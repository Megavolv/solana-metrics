package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Connection struct {
	address string
	log     *log.Logger
}

func NewConnection(address string) *Connection {
	c := &Connection{address: address}
	return c
}

func (c Connection) Health() error {
	resp, err := http.Get(c.address + "/health")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	sb := string(body)
	if sb != "ok" {
		return errors.New("Rpc server is not ready, say: " + sb)
	}

	return nil
}
