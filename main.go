package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

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
	filteredSubmissions := make(map[string]*leetcode.SubmissionDumpNode)
	submissionsVars := make(map[string]interface{})
	submissionsVars["offset"] = 0
	submissionsVars["limit"] = 20
	hasNext := true

	for hasNext {

		var sr SubmissionsListNodeResponse
		content, err := gql.Query(submissionsQuery, submissionsVars)
		if err != nil {
			panic(err)
		}

		if err = json.Unmarshal(content, &sr); err != nil {
			panic(err)
		}

		submissionList := sr.Data.SubmissionList

		if submissionList != nil && submissionList.Submissions != nil {
			for _, submission := range submissionList.Submissions {
				if submission != nil {
					if _, ok := filteredSubmissions[submission.Title]; !ok {

						dirName := fmt.Sprintf("%s-%s", questionsMap[submission.Title].QuestionFrontendID, questionsMap[submission.Title].QuestionTitleSlug)
						dirPath := path.Join(".", dirName)
						info, _ := os.Stat(dirPath)
						if !info.IsDir() {
							os.Mkdir(dirPath, 755)
						}

						filteredSubmissions[submission.Title] = submission
					}
				}
			}
		}

		submissionsVars["offset"] = submissionsVars["offset"].(int) + submissionsVars["limit"].(int)

		hasNext = submissionList.HasNext
	}

}
