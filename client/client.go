package client

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode   int
	ResponseBody string
}

func RequestApi(url string) Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{500, ""}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{500, ""}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{500, ""}
	}

	out := Response{
		StatusCode:   resp.StatusCode,
		ResponseBody: string(body),
	}

	return out
}
