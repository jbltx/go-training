package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	leetcode "./graphql-gen/examples/leetcode-client"
)

const (
	submissionsQuery = `
query SubmissionsQuery($offset: Int, $limit: Int) {
	submissionList(
		offset:$offset
		limit:$limit
	) {
		lastKey
		hasNext    
		submissions {
			title
		}
	}
}
`
	questionsQuery = ``
)

type SubmissionsListNodeResponse struct {
	Data *struct {
		SubmissionList *leetcode.SubmissionListNode `json:"submissionList"`
	} `json:"data"`
	Errors []*leetcode.GraphQLError `json:"errors"`
}

func main() {

	var jwt string
	flag.StringVar(&jwt, "jwt", "", "JSON Web Token")
	flag.Parse()

	gql := leetcode.NewGraphQLClient()

	cookie := &http.Cookie{
		Name:   "LEETCODE_SESSION",
		Value:  jwt,
		Path:   "/",
		Domain: gql.Endpoint.Hostname(),
	}
	gql.AddCookie(cookie)

	submissionsVars := make(map[string]interface{})
	submissionsVars["offset"] = 0
	submissionsVars["limit"] = 20

	hasNext := true

	for hasNext {

		var resBody SubmissionsListNodeResponse
		content, err := gql.Query(submissionsQuery, submissionsVars)
		if err != nil {
			panic(err)
		}

		if err = json.Unmarshal(content, &resBody); err != nil {
			panic(err)
		}

		submissions := resBody.Data.SubmissionList
		decodedLastKey, _ := url.QueryUnescape(submissions.LastKey)
		fmt.Println("Last Key :", decodedLastKey)

		submissionsVars["offset"] = submissionsVars["offset"].(int) + submissionsVars["limit"].(int)

		hasNext = submissions.HasNext
	}

}
