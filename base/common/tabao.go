package common

import (
	"io"
	"net/http"
)

func RequestTaobao() (string, error) {
	resp, err := http.Get("https://www.taobao.com")

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
