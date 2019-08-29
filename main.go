package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// GraphQLClient ...
type GraphQLClient struct {
	httpClient *http.Client
	Endpoint   string
}

// NewGraphQLClient ...
func NewGraphQLClient(endpoint string) *GraphQLClient {
	return &GraphQLClient{
		httpClient: &http.Client{},
		Endpoint:   endpoint,
	}
}

// SetCookie ...
func (gql *GraphQLClient) SetCookie(key string, val string) {

}

/// AUTO-GENERATED

// SubmissionDumpNode ...
type SubmissionDumpNode struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	StatusDisplay string `json:"statusDisplay"`
	Lang          string `json:"lang"`
	Runtime       string `json:"runtime"`
	Time          string `json:"time"`
	Timestamp     string `json:"timestamp"`
	URL           string `json:"url"`
	IsPending     string `json:"isPending"`
	Memory        string `json:"memory"`
}

// SubmissionStatusNode ...
type SubmissionStatusNode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SubmissionListNode ...
type SubmissionListNode struct {
	LastKey     string                `json:"lastKey"`
	HasNext     bool                  `json:"hasNext"`
	Submissions []*SubmissionDumpNode `json:"submissions"`
}

// ID ...
type ID uint64

// ContributorNode ...
type ContributorNode struct {
	Username   string `json:"username"`
	ProfileURL string `json:"profileUrl"`
	AvatarURL  string `json:"avatarUrl"`
}

// CompanyTagNode ...
type CompanyTagNode struct {
	// WIP
}

// TopicTagNode ...
type TopicTagNode struct {
	// WIP
}

// PostNode ...
type PostNode struct {
	// WIP
}

// CommentNode ...
type CommentNode struct {
	// WIP
}

// SolutionArticleNode ...
type SolutionArticleNode struct {
	// WIP
}

// UserNode ...
type UserNode struct {
	// WIP
}

// TopicNode ...
type TopicNode struct {
	ID                   int                  `json:"id"`
	Title                string               `json:"title"`
	ViewCount            int                  `json:"viewCount"`
	Post                 *PostNode            `json:"post"`
	CommentCount         int                  `json:"commentCount"`
	TopLevelCommentCount int                  `json:"topLevelCommentCount"`
	LastComment          *CommentNode         `json:"lastComment"`
	Tags                 []string             `json:"tags"`
	NodebbTid            int                  `json:"nodebbTid"`
	Pinned               bool                 `json:"pinned"`
	PinnedAt             time.Time            `json:"pinnedAt"`
	Article              *SolutionArticleNode `json:"article"`
	Subscribed           bool                 `json:"subscribed"`
	LastActivity         int                  `json:"lastActivity"`
	Authors              []*UserNode          `json:"authors"`
	Category             *DiscussCategoryNode `json:"category"`
	Index                int                  `json:"index"`
}

// UserRatingNode ...
type UserRatingNode struct {
	ID    ID  `json:"id"`
	Score int `json:"score"`
}

// RatingNode ...
type RatingNode struct {
	ID         ID              `json:"id"`
	Count      int             `json:"count"`
	Average    float32         `json:"average"`
	UserRating *UserRatingNode `json:"userRating"`
}

// ArticleNode ...
type ArticleNode struct {
	ID            ID          `json:"id"`
	Title         string      `json:"title"`
	URL           string      `json:"url"`
	Content       string      `json:"content"`
	ContentTypeID ID          `json:"contentTypeId"`
	CanSeeDetail  bool        `json:"canSeeDetail"`
	Rating        *RatingNode `json:"rating"`
	PaidOnly      bool        `json:"paidOnly"`
}

// CodeSnippetNode ...
type CodeSnippetNode struct {
	// WIP
}

// QuestionNode ...
type QuestionNode struct {
	Title                 string             `json:"title"`
	TitleSlug             string             `json:"titleSlug"`
	QuestionID            string             `json:"questionId"`
	QuestionFrontendID    string             `json:"questionFrontendId"`
	QuestionTitle         string             `json:"questionTitle"`
	QuestionTitleSlug     string             `json:"questionTitleSlug"`
	QuestionType          string             `json:"questionType"`
	TranslatedTitle       string             `json:"translatedTitle"`
	Content               string             `json:"content"`
	TranslatedContent     string             `json:"translatedContent"`
	SessionID             string             `json:"sessionId"`
	CategoryTitle         string             `json:"categoryTitle"`
	Difficulty            string             `json:"difficulty"`
	Stats                 string             `json:"stats"`
	Contributors          []*ContributorNode `json:"contributors"`
	CompanyTags           []*CompanyTagNode  `json:"companyTags"`
	CompanyTagStats       string             `json:"companyTagStats"`
	TopicTags             []*TopicTagNode    `json:"topicTags"`
	SimilarQuestions      string             `json:"similarQuestions"`
	MysqlSchemas          []string           `json:"mysqlSchemas"`
	JudgeType             string             `json:"judgeType"`
	CodeDefinition        string             `json:"codeDefinition"`
	SampleTestCase        string             `json:"sampleTestCase"`
	Likes                 int                `json:"likes"`
	Dislikes              int                `json:"dislikes"`
	Article               string             `json:"article"`
	ArticleTopicID        string             `json:"articleTopicId"`
	IsPaidOnly            bool               `json:"isPaidOnly"`
	Status                string             `json:"status"`
	AllowDiscuss          bool               `json:"allowDiscuss"`
	BoundTopicID          int                `json:"boundTopicId"`
	EnableTestMode        bool               `json:"enableTestMode"`
	MetaData              string             `json:"metaData"`
	EnableRunCode         bool               `json:"enableRunCode"`
	EnableSubmit          bool               `json:"enableSubmit"`
	JudgerAvailable       bool               `json:"judgerAvailable"`
	InfoVerified          bool               `json:"infoVerified"`
	EnvInfo               string             `json:"envInfo"`
	RandomQuestionURL     string             `json:"randomQuestionUrl"`
	SubmitURL             string             `json:"submitUrl"`
	InterpretURL          string             `json:"interpretUrl"`
	QuestionDetailURL     string             `json:"questionDetailUrl"`
	URLManager            string             `json:"urlManager"`
	IsLiked               bool               `json:"isLiked"`
	NextChallengePairs    string             `json:"nextChallengePairs"`
	LibraryURL            string             `json:"libraryUrl"`
	Note                  string             `json:"note"`
	Hints                 []string           `json:"hints"`
	Solution              *ArticleNode       `json:"solution"`
	CodeSnippets          []*CodeSnippetNode `json:"codeSnippets"`
	LangToValidPlayground string             `json:"langToValidPlayground"`
	FrequencyTimePeriod   int                `json:"frequencyTimePeriod"`
}

// ItemNode ...
type ItemNode struct {
	ID       ID            `json:"id"`
	Type     string        `json:"type"`
	Title    string        `json:"title"`
	Info     string        `json:"info"`
	Question *QuestionNode `json:"question"`
	// WIP
}

// ChapterNode ...
type ChapterNode struct {
	ID              ID          `json:"id"`
	Title           string      `json:"title"`
	Slug            string      `json:"slug"`
	Description     string      `json:"description"`
	Items           []*ItemNode `json:"items"`
	DescriptionText string      `json:"descriptionText"`
}

// DiscussCategoryNode ...
type DiscussCategoryNode struct {
	ID               ID                     `json:"id"`
	Title            string                 `json:"title"`
	Slug             string                 `json:"slug"`
	Description      string                 `json:"description"`
	Announcement     string                 `json:"announcement"`
	Path             string                 `json:"path"`
	AnonymousEnabled bool                   `json:"anonymousEnabled"`
	Subcategories    []*DiscussCategoryNode `json:"subcategories"`
	RootCategory     *DiscussCategoryNode   `json:"rootCategory"`
	NumTopics        int                    `json:"numTopics"`
	TitleLink        string                 `json:"titleLink"`
	CanModerate      bool                   `json:"canModerate"`
	CanRewardPosts   bool                   `json:"canRewardPosts"`
}

// CardNode ...
type CardNode struct {
	ID                   ID                   `json:"id"`
	Title                string               `json:"title"`
	Slug                 string               `json:"slug"`
	Description          string               `json:"description"`
	Order                int                  `json:"order"`
	Cards                []*CardNode          `json:"cards"`
	Introduction         string               `json:"introduction"`
	CreatedAt            time.Time            `json:"createdAt"`
	LastModified         time.Time            `json:"lastModified"`
	PaidOnly             bool                 `json:"paidOnly"`
	SequentialOnly       bool                 `json:"sequentialOnly"`
	Published            bool                 `json:"published"`
	Img                  string               `json:"img"`
	Chapters             []*ChapterNode       `json:"chapters"`
	Banner               string               `json:"banner"`
	BannerBackground     string               `json:"bannerBackground"`
	Progress             string               `json:"progress"`
	Items                []*ItemNode          `json:"items"`
	IntroText            string               `json:"introText"`
	CategorySlug         string               `json:"categorySlug"`
	CategoryTitle        string               `json:"categoryTitle"`
	IsFavorite           bool                 `json:"isFavorite"`
	NumChapters          int                  `json:"numChapters"`
	NumItems             int                  `json:"numItems"`
	NumUsersStarted      int                  `json:"numUsersStarted"`
	NumUsersCompleted    int                  `json:"numUsersCompleted"`
	Users                string               `json:"users"`
	PrevCompleteLinkInfo string               `json:"prevCompleteLinkInfo"`
	Popularity           int                  `json:"popularity"`
	IsPreview            bool                 `json:"isPreview"`
	IsFeatured           bool                 `json:"isFeatured"`
	DiscussCategory      *DiscussCategoryNode `json:"discussCategory"`
}

// CategoryNode ...
type CategoryNode struct {
	ID          ID          `json:"id"`
	Title       string      `json:"title"`
	Slug        string      `json:"slug"`
	Description string      `json:"description"`
	Order       int         `json:"order"`
	Cards       []*CardNode `json:"cards"`
}

func (gql *GraphQLClient) submissionList(offset int, limit int, lastKey string, questionSlug string) (*SubmissionListNode, error) {
	content := `
{
	submissionList(
		%s
		%s
		%s
		%s) {
		submissions {
			url
			title
			id
			statusDisplay
			timestamp
		}
	}
}
	`
	var requestBody bytes.Buffer
	requestBodyObj := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     req.q,
		Variables: req.vars,
	}

	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, gql.Endpoint, &requestBody)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	r.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := gql.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var buf bytes.Buffer

	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(&buf).Decode(&gr); err != nil {
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("graphql: server returned a non-200 status code: %v", res.StatusCode)
		}
		return nil, errors.Wrap(err, "decoding response")
	}
}

func (gql *GraphQLClient) allQuestions(containMain bool, containExplore bool, containContest bool) []*QuestionNode {

}

/// AUTO-GENERATED - END

const (
	graphqlURL = "https://leetcode.com/graphql"
)

func main() {

	var jwt string
	flag.StringVar(&jwt, "jwt", "", "JSON Web Token")
	flag.Parse()

	gql := NewGraphQLClient(graphqlURL)
	gql.SetCookie("LEETCODE_SESSION", jwt)

}

/*











type Submission struct {
	ID            int64  `json:"id"`
	Lang          string `json:"lang"`
	Time          string `json:"time"`
	Timestamp     int64  `json:"timestamp"`
	StatusDisplay string `json:"status_display"`
	Runtime       string `json:"runtime"`
	URL           string `json:"url"`
	IsPending     string `json:"is_pending"`
	Title         string `json:"title"`
	Memory        string `json:"memory"`
	Code          string `json:"code"`
	CompareResult string `json:"compare_result"`
}

type Question struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	QuestionTitle      string `json:"questionTitle"`
	Content            string `json:"content"`
	TitleSlug          string `json:"titleSlug"`
	Difficulty         string `json:"difficulty"`
}

func (s *Submission) Slug() string {
	return strings.ToLower(strings.ReplaceAll(s.Title, " ", "-"))
}

type SubmissionResponse struct {
	SubmissionsDump []*Submission `json:"submissions_dump"`
	HasNext         bool          `json:"has_next"`
	LastKey         string        `json:"last_key"`
}

type QuestionResponse struct {
	Data struct {
		Question *Question `json:"question"`
	} `json:"data"`
}

type LeetcodeAPIClient struct {
	*http.Client
}

type LeetcodeGraphQLClient struct {
	*http.Client
}

const (
	leetcodeURL    = "https://leetcode.com"
	submissionsURL = "https://leetcode.com/api/submissions"
	graphqlURL     = "https://leetcode.com/graphql"
)

var (
	graphqlClient *LeetcodeGraphQLClient
	apiClient     *LeetcodeAPIClient
)

func (gql *LeetcodeGraphQLClient) getQuestion(titleSlug string) (*Question, error) {

	res, err := gql.Post(graphqlURL, "application/graphql", )
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		dataRaw, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var res QuestionResponse
		err = json.Unmarshal(dataRaw, &res)
		if err != nil {
			return nil, err
		}

		return res.Data.Question, nil
	}

	return nil, errors.New("Status Code is not OK")
}

func (api *LeetcodeAPIClient) getSubmissions() (map[string]*Submission, error) {
	res, err := api.Get(submissionsURL)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		dataRaw, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var data SubmissionResponse
		err = json.Unmarshal(dataRaw, &data)
		if err != nil {
			return nil, err
		}

		cache := make(map[string]*Submission)

		for _, s := range data.SubmissionsDump {
			if _, ok := cache[s.Title]; !ok {
				cache[s.Title] = s
			}
		}

		return cache, nil
	}

	return nil, errors.New("Status Code is not OK")
}

func init() {

	var jwt string
	flag.StringVar(&jwt, "jwt", "", "JSON Web Token")
	flag.Parse()

	jwtCookie := &http.Cookie{
		Name:     "LEETCODE_SESSION",
		Value:    jwt,
		HttpOnly: false,
		Domain:   "leetcode.com",
		Path:     "/",
	}

	apiJar, _ := cookiejar.New(nil)
	cookiesURL, _ := url.Parse(leetcodeURL)
	apiJar.SetCookies(cookiesURL, []*http.Cookie{jwtCookie})

	graphqlClient = &LeetcodeGraphQLClient{
		Client: &http.Client{
			Jar: apiJar,
		},
	}

	apiClient = &LeetcodeAPIClient{
		Client: &http.Client{
			Jar: apiJar,
		},
	}
}

func main() {

	submissions, err := apiClient.getSubmissions()
	if err != nil {
		panic(err)
	}
	for _, v := range submissions {
		question, err := graphqlClient.getQuestion(v.Slug())
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s-%s\n", question.QuestionFrontendID, question.TitleSlug)
	}
}
*/
