package client

type Client struct{ Folder string }

func New(folder string) *Client {
	return &Client{Folder: folder}
}
