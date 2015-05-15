package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func (c *RancherClient) GetHostAccess(resource Resource, action string, input interface{}) (*websocket.Conn, error) {
	var resp HostAccess
	url := resource.Actions[action]
	if url == "" {
		return nil, fmt.Errorf("Failed to find action: %s", action)
	}

	err := c.Post(url, input, &resp)
	if err != nil {
		return nil, err
	}

	url = fmt.Sprintf("%s?token=%s", resp.Url, resp.Token)

	conn, _, err := c.Websocket(url, nil)

	return conn, err
}
