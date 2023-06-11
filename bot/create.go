package bot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/g8rswimmer/go-twitter/v2"
)

func createTweet(client *twitter.Client, text string, dryrun bool) error {
	req := twitter.CreateTweetRequest{
		Text: text,
	}

	fmt.Println(req.Text)
	if dryrun {
		return nil
	}

	tweetResponse, err := client.CreateTweet(context.Background(), req)
	if err != nil {
		return err
	}

	enc, err := json.MarshalIndent(tweetResponse, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(enc))

	return nil
}
