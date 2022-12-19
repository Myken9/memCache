package memcache

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrMalformedKey = errors.New("malformed: key is too long or contains invalid characters")
)

type Item struct {
	Key string
	Val string
}

type Client struct {
	cmd *exec.Cmd
}

func New(s string) *Client {
	cmd := exec.Command("telnet", "localhost", s)
	out := bytes.Buffer{}
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
	return &Client{cmd: cmd}
}

func (c *Client) Get(key string) (*Item, error) {

	return &Item{}, nil
}

func (c *Client) Set(item Item) error {
	if !legalKey(item.Key) {
		return ErrMalformedKey
	}

	return nil
}

func (c *Client) Delete(key string) error {
	return nil
}

func legalKey(key string) bool {
	if len(key) > 250 {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i] <= ' ' || key[i] == 0x7f {
			return false
		}
	}
	return true
}

func useCommand(s string) ([]byte, error) {
	cmd := exec.Command(s)
	return cmd.Output()
}
