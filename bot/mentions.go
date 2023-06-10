package bot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/g8rswimmer/go-twitter/v2"
)

func fetchMentionTimeline(client *twitter.Client, twitterUserID string) error {
	opts := twitter.UserMentionTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldPublicMetrics, twitter.TweetFieldContextAnnotations},
		UserFields:  []twitter.UserField{twitter.UserFieldUserName},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		MaxResults:  5,
	}

	timeline, err := client.UserMentionTimeline(context.Background(), twitterUserID, opts)
	if err != nil {
		return fmt.Errorf("user mention timeline error: %v", err)
	}

	dictionaries := timeline.Raw.TweetDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(enc))

	metaBytes, err := json.MarshalIndent(timeline.Meta, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(metaBytes))

	return nil
}
