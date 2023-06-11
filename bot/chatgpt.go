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
	You will reply to the user as that character.
	All tweets will be from the characters perspective.
	Your only job is to generate a response in the form of a tweet.
	Make sure the tweet is well formed and ready to post.
	Generate exactly one tweet.
	The tweet should be 280 characters or less.
	Do not use your twitter handle in the tweet.
	`
	defaultPrompt = "What's new? Provide an update for us!"
)

func systemRoleContent(handle, characterBackground string) string {
	content := promptPrepend + characterBackground + promptAppend + `Your handle on twitter is @` + handle + `. Do not use @` + handle + ` in the tweet.`

	return content
}

func generateChatGPTTweet(openaiApiKey, characterBackground, twitterHandle, prompt string) (string, error) {
	client := openai.NewClient(openaiApiKey)

	if prompt == "" {
		prompt = defaultPrompt
	}

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
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
