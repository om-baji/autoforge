package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// RenderFromGist downloads a Terraform template from a Gist
// and replaces placeholders using the provided map.
//
// Inputs:
//
//	gistURL      -> raw OR normal gist URL
//	replacements -> map["__PLACEHOLDER__"] = "value"
//
// Output:
//
//	rendered terraform file as string
func RenderFromGist(gistURL string, replacements map[string]string) (string, error) {
	rawURL := normalizeGistURL(gistURL)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to download gist: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch gist, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(body)

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content, nil
}

// normalizeGistURL converts a normal GitHub Gist URL into a raw URL
func normalizeGistURL(url string) string {
	if strings.Contains(url, "raw.githubusercontent.com") {
		return url
	}

	// Example:
	// https://gist.github.com/user/hash
	// -> https://gist.githubusercontent.com/user/hash/raw
	url = strings.Replace(url, "gist.github.com", "gist.githubusercontent.com", 1)

	if !strings.Contains(url, "/raw") {
		url = url + "/raw"
	}

	return url
}
