package common

import (
	"net/http"

	"github.com/shurcooL/graphql"
)

var URL = "https://api.github.com/graphql"

type WithHeader struct {
	http.Header
	rt http.RoundTripper
}

func withHeader(rt http.RoundTripper) WithHeader {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return WithHeader{Header: make(http.Header), rt: rt}
}

func (h WithHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.Header {
		req.Header[k] = v
	}

	return h.rt.RoundTrip(req)
}

func CreateGQLClient(token string) *graphql.Client {
	httpClient := http.DefaultClient

	rt := withHeader(httpClient.Transport)
	rt.Set("GraphQL-Features", "discussions_api")
	rt.Set("Authorization", "Bearer "+token)
	httpClient.Transport = rt

	GQLClient := graphql.NewClient(
		URL,
		httpClient,
	)
	return GQLClient
}
