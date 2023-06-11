package bot

import (
	"fmt"
	"net/http"
)

type authorizer struct {
	Token string
}

func (a authorizer) Add(req *http.Request) {
	if a.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
	}
}
