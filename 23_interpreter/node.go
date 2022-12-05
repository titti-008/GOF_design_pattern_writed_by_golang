package main

type node interface {
	parse(c *context) error
	String() string
}
