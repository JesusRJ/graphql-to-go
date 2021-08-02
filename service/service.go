package service

import (
	"context"
	"graphql-to-go/entity"
	"net/http"
	"time"

	"github.com/machinebox/graphql"
)

const (
	querySchema = `
    query{
        __schema {
          types {
            name
            kind
            fields {
              name
              args{
                  name
                  type{
                    name
                    kind
                  }
                  defaultValue
              }
              type {
                name
                kind
                ofType {
                  name
                  kind
                }
              }
            }
          }
        }
      }
    `

	queryTypeSchema = `
    query($name: String!) {
        __type(name: $name) {
          name
          kind
          fields{
            name
            args{
              name
              type{
                name
              }
              defaultValue
            }
            type{
              name
              kind
              ofType {
                name
                kind
              }
            }
          }
        }
      }
    `
)

// Service Provide methods to get grapql schemas
type Service struct {
	client      *graphql.Client
	accessToken string
}

// NewService Create a new Service instance
func NewService(url, accessToken string, timeout int) *Service {
	httpClient := http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	graphqlClient := graphql.NewClient(url, graphql.WithHTTPClient(&httpClient))

	return &Service{client: graphqlClient, accessToken: accessToken}
}

// NewRequest Create a new graphql.Request with default headers
func (s *Service) NewRequest(q string) (request *graphql.Request) {
	request = graphql.NewRequest(q)

	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Add("Accept", "application/json")
	if s.accessToken != "" {
		request.Header.Set("Authorization", `Bearer `+s.accessToken)
	}

	return request
}

// GetSchema Retrieve a type schema
func (s *Service) GetSchema(ctx context.Context) (*entity.ResponseDataSchema, error) {
	request := s.NewRequest(querySchema)
	responseData := new(entity.ResponseDataSchema)

	if err := s.client.Run(ctx, request, responseData); err != nil {
		return nil, err
	}

	return responseData, nil
}

// GetType Retrieve a type schema
func (s *Service) GetType(ctx context.Context, typeName string) (*entity.ResponseDataType, error) {
	request := s.NewRequest(queryTypeSchema)
	request.Var("name", typeName)

	responseData := new(entity.ResponseDataType)

	if err := s.client.Run(ctx, request, responseData); err != nil {
		return nil, err
	}

	return responseData, nil
}
