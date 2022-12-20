package memcache

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"memcach/pkg/server"
	"net"
	"strconv"
	"strings"
)

var (
	ErrMalformedKey = errors.New("malformed: key is too long or contains invalid characters")
)

type Client struct {
	conn     net.Conn
	buffered bufio.ReadWriter
}

func New(address string) (*Client, error) {
	nc, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return newConnection(nc), nil
}

func newConnection(nc net.Conn) *Client {
	return &Client{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
	}
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		return
	}
	c.conn = nil
}

func (c *Client) Get(key string) (i server.Item, err error) {
	command := "get " + key + "\r\n"
	c.write(command)

	header := c.readLine()
	if strings.HasPrefix(header, "VALUE") {
		chunks := strings.Split(header, " ")
		if len(chunks) < 4 {
			return i, errors.New(fmt.Sprintf("Malformed response: %s", header))
		}
		i.Key = chunks[1]
		size, err := strconv.ParseUint(chunks[3], 10, 64)
		if err != nil {
			return i, errors.New(fmt.Sprintf("%s", err))
		}
		i.Val = c.read(int(size) + 2)[:size]
		header = c.readLine()
	} else if strings.HasPrefix(header, "END") {
		return i, errors.New("memcached: cache miss")
	}
	return
}

func (c *Client) Set(item server.Item) (err error) {
	if !legalKey(item.Key) {
		return ErrMalformedKey
	}

	valLen := strconv.Itoa(len([]byte(item.Val)))
	command := "set " + item.Key + " 0 900 " + valLen + "\r\n"
	c.write(command)
	c.write(item.Val)
	c.write("\r\n")
	reply := c.readLine()
	if strings.Contains(reply, "ERROR") {
		return errors.New(fmt.Sprintf("Malformed response: %s", reply))
	}
	return
}

func (c *Client) Delete(key string) error {
	command := "delete " + key + "\r\n"
	c.write(command)
	reply := c.readLine()
	if strings.Contains(reply, "ERROR") {
		panic(reply)
	}
	return nil
}

func (c *Client) write(s string) {
	_, err := c.buffered.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}

func (c *Client) readLine() string {
	c.flush()
	l, isPrefix, err := c.buffered.ReadLine()
	if isPrefix || err != nil {
		panic(errors.New(fmt.Sprintf("Prefix: %v, %s", isPrefix, err)))
	}
	return string(l)
}

func (c *Client) read(count int) string {
	c.flush()
	b := make([]byte, count)
	if _, err := io.ReadFull(c.buffered, b); err != nil {
		panic(err)
	}
	return string(b)
}

func (c *Client) flush() {
	if err := c.buffered.Flush(); err != nil {
		panic(err)
	}
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
