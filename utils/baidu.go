package utils

import (
	"io"
	"net/http"
)

func RequestBaidu() (string, error) {
	resp, err := http.Get("https://www.baidu.com")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil

}
