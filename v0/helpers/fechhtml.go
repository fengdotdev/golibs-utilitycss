package helpers

import (
	"fmt"
	"io"
	"net/http"
)


// FetchHTML retrieves the HTML content from the specified URL.
// It returns the HTML as a string or an error if the request fails.
func FetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch HTML: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
