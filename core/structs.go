package core

import (
	"fmt"
	"net/http"
)

type Curl struct {
	Method string `json:"method"`
	Url string `json:"url"`
	Headers http.Header
	Body string `json:"body"`
	Error string `json:"error"`
}

func (c *Curl) Print() string {
	return fmt.Sprintf( "[%s] 'uri:%s' 'body:%s'", c.Method, c.Url, c.Body)
}


