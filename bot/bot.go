package bot

import (
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
)

const (
	twitterAPIHost = "https://api.twitter.com"
)

func createTwitterAppClient(token string) *twitter.Client {
	return &twitter.Client{
		Authorizer: authorizer{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   twitterAPIHost,
	}
}

func createTwitterOAuthClient(consumerKey, consumerSecret, accessToken, accessSecret string) *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return &twitter.Client{
		Authorizer: &authorizer{},
		Client:     httpClient,
		Host:       twitterAPIHost,
	}
}

func Run(twitterUserID, twitterHandle, twitterBearerToken, twitterApiKey, twitterApiSecret, twitterAccessToken, twitterAccessSecret, openaiApiKey, characterBackground, prompt string, dryrun, bot bool) error {
	// clientApp := createTwitterAppClient(twitterBearerToken)
	clientOAuth := createTwitterOAuthClient(twitterApiKey, twitterApiSecret, twitterAccessToken, twitterAccessSecret)

	if bot {
		err := fetchMentionTimeline(clientOAuth, twitterUserID)
		if err != nil {
			return err
		}

		return nil
	}

	tweet, err := generateChatGPTTweet(openaiApiKey, characterBackground, twitterHandle, prompt)
	if err != nil {
		return err
	}

	err = createTweet(clientOAuth, tweet, dryrun)
	if err != nil {
		return err
	}

	return nil
}
