package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type context struct {
	tokens    []string
	lastToken string
	index     int
	window    *window
}

func newContext(text string, w *window) *context {
	c := new(context)

	tokens := strings.Split(text, " ")
	fmt.Println(tokens)

	c.tokens = tokens
	c.window = w
	c.index = 0
	c.nextToken()
	return c
}

func (c *context) nextToken() string {
	if c.index < len(c.tokens) {
		c.lastToken = c.tokens[c.index]
		c.index = c.index + 1
	} else {
		c.lastToken = ""
	}
	return c.lastToken
}

func (c *context) currentToken() string {
	return c.lastToken
}

func (c *context) skipToken(token string) error {
	if c.currentToken() == "" {
		return fmt.Errorf("Error: %v is expected, but no more token is found.\n", token)
	} else if token != c.currentToken() {
		return fmt.Errorf("Error: %v is expected, but %v is found.\n", token, c.currentToken())
	}

	c.nextToken()

	return nil
}

func (c *context) currentNumber() (int, error) {
	if c.currentToken() == "" {
		return 0, errors.New("Error: No more token.")
	}
	number, err := strconv.Atoi(c.currentToken())
	if err != nil {
		return 0, err
	}

	return number, nil
}
