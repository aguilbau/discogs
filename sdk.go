package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	BaseURL = "https://api.discogs.com"
)

type ClientConfig struct {
	Token          string
	ConsumerKey    string
	ConsumerSecret string
}

func (c ClientConfig) getAuthorization() (string, error) {
	if c.Token != "" {
		return fmt.Sprintf("Discogs token=%s", c.Token), nil
	}
	if c.ConsumerKey == "" || c.ConsumerSecret == "" {
		return "", errors.New("you must specify a token or (ConsumerKey and ConsumerSecret)")
	}
	return fmt.Sprintf(
			"Discogs key=%s, secret=%s",
			c.ConsumerKey,
			c.ConsumerSecret),
		nil
}

type Client struct {
	auth string
}

func NewClient(config ClientConfig) (*Client, error) {
	auth, err := config.getAuthorization()
	if err != nil {
		return nil, err
	}
	c := Client{
		auth: auth,
	}

	return &c, nil
}

func (c *Client) doRequest(
	ctx context.Context,
	verb, path string,
	body interface{},
	result interface{}) error {

	var b io.ReadWriter

	if body != nil {
		b = new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequestWithContext(ctx, verb, BaseURL+path, b)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", c.auth)
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// todo : handle errors
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) GetRelease(
	ctx context.Context,
	releaseID int64,
	currency *Currency) (*Release, error) {

	path := "/releases/" + strconv.FormatInt(releaseID, 10)
	if currency != nil {
		path += "?" + url.QueryEscape(string(*currency))
	}

	var release Release
	if err := c.doRequest(ctx, "GET", path, nil, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

func (c *Client) GetLabelReleases(
	ctx context.Context,
	labelID int64,
	pagination *PaginationParams) (*LabelReleases, error) {

	path := fmt.Sprintf("/labels/%d/releases", labelID)
	if pagination != nil {
		path += "?" + pagination.toQuery()
	}

	var labelReleases LabelReleases
	if err := c.doRequest(ctx, "GET", path, nil, &labelReleases); err != nil {
		return nil, err
	}
	return &labelReleases, nil
}
