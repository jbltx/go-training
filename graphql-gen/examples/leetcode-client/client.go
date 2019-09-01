package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

/* The file is auto-generated, do not edit this ! */

const (
	endpointStr = "https://leetcode.com/graphql"
	contentType = "application/json; charset=utf-8"
)

// GraphQLClient ...
type GraphQLClient struct {
	httpClient *http.Client
	Endpoint   *url.URL
	Cookies    []*http.Cookie
}

// GraphQLQuery ...
type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// GraphQLResponse ...
type GraphQLResponse struct {
	Data   interface{}     `json:"data"`
	Errors []*GraphQLError `json:"errors"`
}

// GraphQLError ...
type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
}

// NewGraphQLClient ...
func NewGraphQLClient() *GraphQLClient {

	url, _ := url.Parse(endpointStr)
	jar, _ := cookiejar.New(nil)

	return &GraphQLClient{
		httpClient: &http.Client{
			Jar: jar,
		},
		Endpoint: url,
		Cookies:  make([]*http.Cookie, 0),
	}
}

// AddCookie ...
func (gql *GraphQLClient) AddCookie(c *http.Cookie) {
	cookies := gql.httpClient.Jar.Cookies(gql.Endpoint)
	cookies = append(cookies, c)
	gql.httpClient.Jar.SetCookies(gql.Endpoint, cookies)
}

// Query ...
func (gql *GraphQLClient) Query(query string, vars map[string]interface{}) ([]byte, error) {
	var requestBody bytes.Buffer

	requestBodyObj := GraphQLQuery{
		Query:     query,
		Variables: vars,
	}

	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		return nil, err
	}

	res, err := gql.httpClient.Post(gql.Endpoint.String(), contentType, &requestBody)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Error Response Status : " + res.Status)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// Get is the legacy HTTP GET query
func (gql *GraphQLClient) Get(uri string, params map[string]string) ([]byte, error) {

	if params != nil && len(params) > 0 {
		pStr := "?"
		for k, v := range params {
			if len(pStr) > 1 {
				pStr += "&"
			}
			pStr += fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v))
		}

		uri += pStr
	}

	fmt.Println("GET", uri)
	res, err := gql.httpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Error Response Status : " + res.Status)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

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

// SubmissionDumpNodeLegacy ...
type SubmissionDumpNodeLegacy struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	StatusDisplay string `json:"status_display"`
	Lang          string `json:"lang"`
	Runtime       string `json:"runtime"`
	Time          string `json:"time"`
	Timestamp     int    `json:"timestamp"`
	URL           string `json:"url"`
	IsPending     string `json:"is_pending"`
	Memory        string `json:"memory"`
	Code          string `json:"code"`
	CompareResult string `json:"compare_result"`
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

type SubmissionListNodeLegacy struct {
	LastKey     string                      `json:"last_key"`
	HasNext     bool                        `json:"has_next"`
	Submissions []*SubmissionDumpNodeLegacy `json:"submissions_dump"`
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
