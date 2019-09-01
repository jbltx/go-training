package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	leetcode "github.com/jbltx/go-training/graphql-gen/examples/leetcode-client"
)

const (
	submissionsQuery = `
query SubmissionsQuery($offset: Int, $limit: Int) {
	submissionList(
		offset:$offset
		limit:$limit
	) {
		# lastKey
		hasNext    
		submissions {
			title
		}
	}
}
`
	questionsQuery = `
query {
	allQuestions(containMain: true) {
		questionFrontendId
		questionTitle
		questionTitleSlug
	}
}	
`
)

type SubmissionsListNodeResponse struct {
	Data *struct {
		SubmissionList *leetcode.SubmissionListNode `json:"submissionList"`
	} `json:"data"`
	Errors []*leetcode.GraphQLError `json:"errors"`
}

type AllQuestionsResponse struct {
	Data *struct {
		AllQuestions []*leetcode.QuestionNode `json:"allQuestions"`
	} `json:"data"`
	Errors []*leetcode.GraphQLError `json:"errors"`
}

func main() {

	// parse arguments from command-line
	var jwt string
	flag.StringVar(&jwt, "jwt", "", "JSON Web Token")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// configure GraphQL client
	gql := leetcode.NewGraphQLClient()
	cookie := &http.Cookie{
		Name:   "LEETCODE_SESSION",
		Value:  jwt,
		Path:   "/",
		Domain: gql.Endpoint.Hostname(),
	}
	gql.AddCookie(cookie)

	// get all available questions first
	rawResponse, err := gql.Query(questionsQuery, nil)
	if err != nil {
		panic(err)
	}
	var res AllQuestionsResponse
	if err = json.Unmarshal(rawResponse, &res); err != nil {
		panic(err)
	}

	// put the list of questions in a hashmap for convenience
	questionsMap := make(map[string]*leetcode.QuestionNode)
	if res.Data.AllQuestions != nil {
		for _, question := range res.Data.AllQuestions {
			questionsMap[question.QuestionTitle] = question
		}
	}

	// get all submissions, and filter it by storing most recent ones in a hash map
	filteredSubmissions := make(map[string]*leetcode.SubmissionDumpNodeLegacy)
	submissionsVars := make(map[string]string)
	submissionsOffset := 0
	submissionsLimit := 20
	submissionsVars["offset"] = strconv.Itoa(submissionsOffset)
	submissionsVars["limit"] = strconv.Itoa(submissionsLimit)
	hasNext := true

	for hasNext {

		var submissionList leetcode.SubmissionListNodeLegacy
		time.Sleep(1 * time.Second) // little hack to not get banned
		// For now I use the legacy API for submissions, because I didn't find a way to get the submitted
		// code using the SubmissionDumpNode of GraphQL API... I suspect Leetcode to not using GraphQL that much :S
		content, err := gql.Get("https://leetcode.com/api/submissions/", submissionsVars)
		if err != nil {
			panic(err)
		}

		if err = json.Unmarshal(content, &submissionList); err != nil {
			panic(err)
		}

		if submissionList.Submissions != nil {
			for _, submission := range submissionList.Submissions {
				if submission != nil && submission.Lang == "golang" {
					if _, ok := filteredSubmissions[submission.Title]; !ok {

						// create the directory for the submission
						dirName := fmt.Sprintf("%s-%s", questionsMap[submission.Title].QuestionFrontendID, questionsMap[submission.Title].QuestionTitleSlug)
						dirPath := path.Join(cwd, dirName)
						info, err := os.Stat(dirPath)
						if err != nil || !info.IsDir() {
							os.Mkdir(dirPath, 755)
						}

						// write submission code in a go file
						codeFilepath := path.Join(dirPath, questionsMap[submission.Title].QuestionTitleSlug+".go")
						f, err := os.Create(codeFilepath)
						if err != nil {
							panic(err)
						}
						f.WriteString(fmt.Sprintf("package main\n\n%s", submission.Code))
						if err = f.Close(); err != nil {
							panic(err)
						}

						// store the submission in the map
						filteredSubmissions[submission.Title] = submission
					}
				}
			}
		}

		submissionsOffset += submissionsLimit
		submissionsVars["offset"] = strconv.Itoa(submissionsOffset)

		hasNext = submissionList.HasNext
	}

}
