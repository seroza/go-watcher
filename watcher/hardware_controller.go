package main

import (
	"fmt"
	"log"
	"net/http"
)

//Command for watcher control
type Command struct {
	Command string
	Value   string
}

func (cmd *Command) parseReq(r *http.Request) (err error) {
	if err = r.ParseForm(); err != nil {
		log.Print("ParseForm error!")
		return
	}

	cmd.Command = r.PostForm["commandSelect"][0]
	cmd.Value = r.PostForm["rigSelect"][0]
	return
}

//NewCommand is a short way to create initialized Command
func NewCommand(r *http.Request) (Command, error) {
	var cmd Command
	err := cmd.parseReq(r)
	return cmd, err
}

func (cmd Command) String() string {
	return fmt.Sprintf("%s:%s", cmd.Command, cmd.Value)
}
