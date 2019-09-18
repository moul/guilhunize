package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"moul.io/guilhunize/guilhunize"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch {
	case request.HTTPMethod == "GET" && request.Path == "/api/quote":
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       guilhunize.Quote(),
			Headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
		}, nil
	case request.HTTPMethod == "POST" && request.Path == "/api/guilhunize":
		if request.Body == "" {
			return nil, fmt.Errorf("missing body; you can use something like:\n\necho refactor | http POST https://guilhunize.moul.io/api/guilhunize --body")
		}
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       guilhunize.Guilhunize(request.Body),
			Headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
		}, nil
	default:
		out, _ := json.MarshalIndent(request, "", "  ")
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("error: unknown method/path.\n\n%s", string(out)),
			Headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
