package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// GraphQLObjectType ...
type GraphQLObjectType struct {
	Name string `json:"name"`
}

// GraphQLNamedType ...
type GraphQLNamedType interface {
}

// GraphQLSchemaSettings ...
type GraphQLSchemaSettings struct {
	Query        *GraphQLObjectType
	Mutation     *GraphQLObjectType
	Subscription *GraphQLObjectType
	Types        []*GraphQLNamedType
	Directives   []*Directive
}

// GraphQLScalarType ...
type GraphQLScalarType struct {
	Name        string
	Description string
}

type TypeRef struct {
	Name   string   `json:"name"`
	Kind   string   `json:"kind"`
	OfType *TypeRef `json:"ofType"`
}

type FullType struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Fields      []struct {
		Name              string        `json:"name"`
		Description       string        `json:"description"`
		Args              []*InputValue `json:"args"`
		Type              *TypeRef      `json:"type"`
		IsDeprecated      bool          `json:"isDeprecated"`
		DeprecationReason string        `json:"deprecationReason"`
	}
	InputFields []*InputValue `json:"inputFields"`
	Interfaces  []*TypeRef    `json:"interfaces"`
	EnumValues  []struct {
		Name              string `json:"name"`
		Description       string `json:"description"`
		IsDeprecated      bool   `json:"isDeprecated"`
		DeprecationReason string `json:"deprecationReason"`
	}
	PossibleTypes []*TypeRef `json:"possibleTypes"`
}

type InputValue struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        *TypeRef `json:"type"`
}

type Directive struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Locations   []string      `json:"locations"`
	Args        []*InputValue `json:"args"`
}

type Schema struct {
	QueryType        *GraphQLObjectType `json:"queryType"`
	MutationType     *GraphQLObjectType `json:"mutationType"`
	SubscriptionType *GraphQLObjectType `json:"subscriptionType"`
	Types            []*FullType        `json:"types"`
	Directives       []*Directive       `json:"directives"`
}

type IntrospectionResponse struct {
	Data *struct {
		Schema *Schema `json:"__schema"`
	} `json:"data"`
}

type GraphQLSchema struct {
}

const (
	introspectionQuery = `
query IntrospectionQuery {
	__schema {
		queryType {
		name
		}
		mutationType {
		name
		}
		subscriptionType {
		name
		}
		types {
		...FullType
		}
		directives {
		name
		description
		locations
		args {
			...InputValue
		}
		}
	}
}
	
fragment FullType on __Type {
	kind
	name
	description
	fields(includeDeprecated: true) {
		name
		description
		args {
		...InputValue
		}
		type {
		...TypeRef
		}
		isDeprecated
		deprecationReason
	}
	inputFields {
		...InputValue
	}
	interfaces {
		...TypeRef
	}
	enumValues(includeDeprecated: true) {
		name
		description
		isDeprecated
		deprecationReason
	}
	possibleTypes {
		...TypeRef
	}
}
	
fragment InputValue on __InputValue {
	name
	description
	type {
		...TypeRef
	}
}
	
fragment TypeRef on __Type {
	kind
	name
	ofType {
		kind
		name
		ofType {
		kind
		name
		ofType {
			kind
			name
			ofType {
			kind
			name
			ofType {
				kind
				name
				ofType {
				kind
				name
				ofType {
					kind
					name
				}
				}
			}
			}
		}
		}
	}
}
	  
`

	INTROSPECTION_FAILED_CODE = 1
	STATS_FAILED_CODE         = 2
	SCHEMA_GEN_FAILED_CODE    = 3
	CLIENT_GEN_FAILED_CODE    = 4
	NO_ENDPOINT_CODE          = 5
)

func getIntrospectionResponse(endpoint string) (*IntrospectionResponse, error) {
	var requestBody bytes.Buffer

	requestBodyObj := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     introspectionQuery,
		Variables: nil,
	}

	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", endpoint, &requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Error Response Status : " + res.Status)
	}

	var resBody IntrospectionResponse

	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	return &resBody, nil
}

func showStatistics(s Schema) error {
	// TODO : Show stats

	return nil
}

func buildType(t *FullType) interface{} {
	return nil
}

func getObjectType(s interface{}) {

}

func buildDirective(d *Directive) {

}

func objectValues(map[string]interface{}) []*GraphQLNamedType {
	return nil
}

func newGraphQLSchema(cfg *GraphQLSchemaSettings) *GraphQLSchema {
	return nil
}

var specifiedScalarTypes []*GraphQLScalarType
var introspectionTypes []*GraphQLScalarType

func generateGQLSchema(schemaIntrospection Schema) *GraphQLSchema {
	// TODO : Generate GraphQL schema using introspection response

	typeMap := make(map[string]interface{})
	for _, t := range schemaIntrospection.Types {
		typeMap[t.Name] = buildType(t)
	}

	for _, t := range append(specifiedScalarTypes, introspectionTypes...) {
		if _, ok := typeMap[t.Name]; ok {
			typeMap[t.Name] = t
		}
	}

	queryType := schemaIntrospection.QueryType
	if queryType != nil {
		getObjectType(schemaIntrospection.QueryType)
	}

	mutationType := schemaIntrospection.MutationType
	if mutationType != nil {
		getObjectType(schemaIntrospection.MutationType)
	}

	subscriptionType := schemaIntrospection.SubscriptionType
	if subscriptionType != nil {
		getObjectType(schemaIntrospection.SubscriptionType)
	}

	directives := schemaIntrospection.Directives
	if directives != nil {
		for _, d := range directives {
			buildDirective(d)
		}
	} else {
		directives = []*Directive{}
	}

	return newGraphQLSchema(&GraphQLSchemaSettings{
		Query:        queryType,
		Mutation:     mutationType,
		Subscription: subscriptionType,
		Types:        objectValues(typeMap),
		Directives:   directives,
	})
}

func generateGolangClient(s Schema, gql GraphQLSchema) error {
	// TODO : Generate client using the schema

	return nil
}

func main() {

	var genClient bool
	var genSchema bool
	var showStats bool
	var endpoint string
	flag.BoolVar(&genClient, "client", false, "Generate client in Golang")
	flag.BoolVar(&genSchema, "schema", false, "Generate GraphQL schema")
	flag.BoolVar(&showStats, "stats", false, "Show statistics about the API")
	flag.StringVar(&endpoint, "endpoint", "https://leetcode.com/graphql", "The GraphQL API endpoint")
	flag.Parse()

	if !showStats && !genSchema && !genClient {
		fmt.Println("Nothing to do, please provide one of the available options flags")
		os.Exit(0)
	}

	if endpoint == "" {
		fmt.Println("Please provide a valid endpoint in the command-line")
		os.Exit(NO_ENDPOINT_CODE)
	}

	fmt.Println("Introspecting...")

	introspection, err := getIntrospectionResponse(endpoint)

	if err != nil {
		fmt.Printf("Introspection of the API has failed : %v", err)
		os.Exit(INTROSPECTION_FAILED_CODE)
	}

	if showStats {
		err = showStatistics(*introspection.Data.Schema)
		if err != nil {
			fmt.Printf("Statistics dislay has failed : %v", err)
			os.Exit(STATS_FAILED_CODE)
		}
	}

	var gqlSchema *GraphQLSchema

	if genSchema || genClient {
		gqlSchema = generateGQLSchema(*introspection.Data.Schema)
		if err != nil {
			fmt.Printf("GraphQL schema generation has failed : %v", err)
			os.Exit(SCHEMA_GEN_FAILED_CODE)
		}
	}

	if genClient {
		generateGolangClient(*introspection.Data.Schema, *gqlSchema)
		if err != nil {
			fmt.Printf("Golang client generation has failed : %v", err)
			os.Exit(CLIENT_GEN_FAILED_CODE)
		}
	}

	os.Exit(0)
}
