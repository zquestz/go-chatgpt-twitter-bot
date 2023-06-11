package bot

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

const (
	promptPrepend = `All the content below up to !-----! is a character description.
	`
	promptAppend = `!-----!
	You are a tweet generator for the character described above.
	You will become that character.
	All tweets will be from the characters perspective.
	Make sure the content is interesting and unique.
	Your only job is to generate tweets.
	Make sure all tweets are well formed and ready to post.
	You will generate exactly one tweet.
	Do not use your twitter handle in the tweet.
	`
)

func systemRoleContent(handle, characterBackground string) string {
	content := promptPrepend + characterBackground + promptAppend + `Your handle on twitter is @` + handle + `. Do not use @` + handle + ` in the tweet.`

	return content
}

func generateChatGPTTweet(openaiApiKey, characterBackground, twitterHandle string) (string, error) {
	client := openai.NewClient(openaiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemRoleContent(twitterHandle, characterBackground),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a tweet!",
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
