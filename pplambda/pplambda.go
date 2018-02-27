package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/brnsampson/go-partyparrot/partyparrot"
)

var (
	tokens       = strings.Split(os.Getenv("SLACK_TEAM_TOKENS"), ":")
	errorBackend = errors.New("something went wrong")
)

// Response replaces the default ApiGatewayProxyResponse because that one doesn't handle json strings correctly.
type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// partyHandler handles a request, generates the emoji strings one rune at a time, and replies
func partyHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// decode json
	p, err := url.ParseQuery(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, fmt.Errorf("error parsing request body %q: %s", request.Body, err)
	}
	fmt.Printf("payload: %s", p)

	// This stuff is for a json payload. However, we have a URL encoded payload.
	//var payload Request
	//err = json.Unmarshal([]byte(request.Body), &payload)
	//if err != nil {
	//	return events.APIGatewayProxyResponse{StatusCode: 500}, fmt.Errorf("error unmarshalling %q: %s", request.Body, err)
	//}

	// Validate.
	isValid := false
	for _, val := range tokens {
		if val == p["token"][0] {
			isValid = true
			break
		}
	}

	if !isValid {
		return events.APIGatewayProxyResponse{StatusCode: 401}, fmt.Errorf("invalid slack token received: %q", p["token"])
	}

	// call partyparrots library
	var b strings.Builder
	_, err = b.WriteString(".\n")
	for _, x := range p["text"][0] {
		_, err := b.WriteString(fmt.Sprintf("%s\n\n\n", partyparrot.PartyConv(x)))
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 404}, err
		}
	}

	// call jsonify
	r := Response{Text: b.String(), ResponseType: "in_channel"}

	rj, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("error formatting response as json %q", r)
	}
	return events.APIGatewayProxyResponse{Body: string(rj), StatusCode: 200}, nil
}

func main() {
	lambda.Start(partyHandler)
}
