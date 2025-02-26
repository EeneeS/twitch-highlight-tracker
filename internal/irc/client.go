package irc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
  conn *net.Conn
  reader *bufio.Reader
}

func NewClient(server string) (*Client, error) {
  conn, err := net.Dial("tcp", server)
  if err != nil {
    fmt.Println("failed connection", err)
    return nil, err
  }
  return &Client{
    conn: &conn,
    reader: bufio.NewReader(conn),
  }, nil
}

func (c *Client) Logon() error {
  err := c.SendData("PASS oath:justinfan1234")
  err = c.SendData("NICK justinfan1234")
  return err
}

func (c *Client) SendData(message string) error {
  _, err := fmt.Fprintf(*c.conn, "%s\r\n", message)
  return err
}

func (c *Client) JoinChannel(channel string) error {
  joinString := fmt.Sprintf("JOIN #%s", channel)
  err := c.SendData(joinString)
  return err
}

func (c *Client) ReadMessages() {
  reader := bufio.NewReader(*c.conn)
  for {
    message, err := reader.ReadString('\n')

    if err != nil {
      fmt.Println("Connection closed", err)
      return
    }

    if strings.Contains(message, "PRIVMSG") {
      parts := strings.Split(message, " :")
      if len(parts) > 1 {
        userMessage := parts[1]
        fmt.Println(userMessage)
      }
    }

    if strings.HasPrefix(message, "PING") {
      c.SendData(strings.Replace(message, "PING", "PONG", 1))
    }
  }
}

func (c *Client) ReadMessage() (string, error) {
  return c.reader.ReadString('\n')
}

func (c *Client) Close() error {
  return (*c.conn).Close()
}
