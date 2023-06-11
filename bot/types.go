package bot

import (
	"fmt"
	"net/http"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

type authorizer struct {
}

func (a *authorizer) Add(req *http.Request) {}
