package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-watcher/watcher/gpio"
)

// Expected commands list
const (
	reboot      = "reboot"
	turnOnOff   = "turn_on_off"
	hardTurnOff = "hard_turn_off"
)

//Command for watcher control
type Command struct {
	Command string
	Value   string
}

func (cmd *Command) execute() {
	switch cmd.Command {
	case reboot:
		gpio.DeviceInstance().Reboot(cmd.Value)
	case turnOnOff:
		gpio.DeviceInstance().TurnOnOff(cmd.Value)
	case hardTurnOff:
		gpio.DeviceInstance().HardTurnOff(cmd.Value)
	default:
		log.Println("wrong command received")
	}
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
