package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

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
	QueryType struct {
		Name string `json:"name"`
	} `json:"queryType"`
	MutationType struct {
		Name string `json:"name"`
	} `json:"mutationType"`
	SubscriptionType struct {
		Name string `json:"name"`
	} `json:"subscriptionType"`
	Types      []*FullType  `json:"types"`
	Directives []*Directive `json:"directives"`
}

type IntrospectionResponse struct {
	Data struct {
		Schema *Schema `json:"__schema"`
	} `json:"data"`
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
)

func getIntrospectionResponse() (*IntrospectionResponse, error) {
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

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", &requestBody)
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

func main() {

	_, err := getIntrospectionResponse()

	if err != nil {
		panic(err)
	}

	// TODO : Generate GraphQL schema using introspection response
	// TODO : Generate client using the schema
}
