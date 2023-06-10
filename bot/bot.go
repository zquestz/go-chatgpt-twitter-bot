package bot

import (
	"net/http"

	"github.com/g8rswimmer/go-twitter/v2"
)

func createTwitterClient(token string) *twitter.Client {
	return &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
}

func Run(twitterUserID string, twitterBearerToken string) error {
	client := createTwitterClient(twitterBearerToken)

	err := fetchMentionTimeline(client, twitterUserID)
	if err != nil {
		return err
	}

	// Make this a call to OpenAI.
	chatGPTTweet := "Tweet text."

	err = createTweet(client, chatGPTTweet)
	if err != nil {
		return err
	}

	return nil
}
