package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getEnv(key string, fallback string, required bool) string {
	value, exists := os.LookupEnv(key)
	if !exists && required {
		fmt.Println(key + " is not set")
		os.Exit(1)
	} else if !exists {
		return fallback
	}
	return value
}

func notify(data url.Values, token string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(data.Encode()))

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func main() {
	accessToken := getEnv("INPUT_ACCESS_TOKEN", "", true)
	message := getEnv("INPUT_MESSAGE", "hello world", false)
	notification := getEnv("INPUT_DISABLE_NOTIFICATION", "false", false)

	data := url.Values{
		"message":              {message},
		"notificationDisabled": {notification},
	}

	resp := notify(data, accessToken)
	fmt.Println(resp.StatusCode)
}
