package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Command for watcher control
type Command struct {
	Command string
	Value   string
}

func (cmd *Command) execute() {
	log.Println("command execution is not implemented yet")
	time.Sleep(time.Second * 4)
}

//ParseReq makes Command struct from reques data
func (cmd *Command) ParseReq(r *http.Request) (err error) {
	if err = r.ParseForm(); err != nil {
		log.Print("ParseForm error!")
		return
	}

	commandSelectField, comOk := r.PostForm["commandSelect"]
	rigSelectField, rigOk := r.PostForm["rigSelect"]

	dataIsValid := comOk && rigOk && len(commandSelectField) == 1 && len(rigSelectField) == 1

	if dataIsValid {
		cmd.Command = commandSelectField[0]
		cmd.Value = rigSelectField[0]
	} else {
		err = errors.New("bad command data")
	}

	return
}

//NewCommand is a short way to create initialized Command
func NewCommand(r *http.Request) (Command, error) {
	var cmd Command
	err := cmd.ParseReq(r)
	return cmd, err
}

func (cmd Command) String() string {
	return fmt.Sprintf("%s:%s", cmd.Command, cmd.Value)
}
