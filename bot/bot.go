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
		Authorizer: authorize{
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

func Run(twitterUserID, twitterHandle, twitterBearerToken, twitterApiKey, twitterApiSecret, twitterAccessToken, twitterAccessSecret, openaiApiKey, openaiPrompt string, dryrun bool) error {
	// clientApp := createTwitterAppClient(twitterBearerToken)
	clientOAuth := createTwitterOAuthClient(twitterApiKey, twitterApiSecret, twitterAccessToken, twitterAccessSecret)

	// err := fetchMentionTimeline(clientOAuth, twitterUserID)
	// if err != nil {
	// 	return err
	// }

	tweet, err := generateChatGPTTweet(openaiApiKey, openaiPrompt, twitterHandle)
	if err != nil {
		return err
	}

	err = createTweet(clientOAuth, tweet, dryrun)
	if err != nil {
		return err
	}

	return nil
}
