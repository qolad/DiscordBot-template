package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"net/http"
	//"context"
)

type AuthTransport struct {
	Transport http.RoundTripper

	Token string //$env:BotToken=""
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqCopy := req.Clone(req.Context())
	reqCopy.Header.Set("Authorization", "Bot "+t.Token)
	reqCopy.Header.Add("User-Agent", "BotInterface/0.1")

	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	return transport.RoundTrip(reqCopy)
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Bot      bool   `json:"bot"`
}

type DiscordClient struct {
	Client *http.Client
}

func NewDiscordClient(token string) *DiscordClient {

	return &DiscordClient{
		Client: &http.Client{
			Transport: &AuthTransport{
				Token: token,
			},
		},
	}

}

func (d *DiscordClient) doRequest(method string, endpoint string) ([]byte, error) {
	req, err := http.NewRequest(method, "https://discord.com/api/v10"+endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("discord returned %s: %s", resp.Status, body)
	}

	return body, nil
}

func (d *DiscordClient) GetCurrentUser() (*User, error) {
	body, err := d.doRequest("GET", "/users/@me")
	if err != nil {
		return nil, err
	}

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil

}

// func (d *DiscordClient) GetServer(){
// 	server, err := do.Request("GET", )
// }

//func (d *DiscordClient) GetServerPreview

// func (d *DiscordClient) CreateMessage() error {
// 	message, err := d.doRequest("POST", "1392903948276338791/messages")
// 	if err != nil {
// 		return nil, err
// 	}
// }

func main() {
	fmt.Println("Live..")

	//issue with environment variable
	/*
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
	*/

  //test Bot / token
	token := os.Getenv("MeowcpyToken")
	if token == "" {
		fmt.Println("Bot Token is not set")
		fmt.Println("run $env:BotToken=\"token_value\"")
		return
	}

	discord := NewDiscordClient(token)
	//discord.GetCurrentUser()
	user, err := discord.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Username)


	if user.Username != ""{

	}
	// fmt.Println("ID:", user.ID)
	// fmt.Println("Username:", user.Username)
	// fmt.Println("Bot:", user.Bot)

}
