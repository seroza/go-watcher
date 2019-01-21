package main

import (
	"log"
)

//CommandsQueue is a concurrent FIFO container for commands
type CommandsQueue struct {
	ch chan Command
}

//NewCommandQueue returns initialized CommandQueue ptr
func NewCommandQueue() (cq *CommandsQueue) {
	cq = &CommandsQueue{make(chan Command)}
	return cq
}

//PushBack is just PushBack
func (q *CommandsQueue) PushBack(cmd *Command) {
	q.ch <- *cmd
}

//PopFront allows to get next element from queue
func (q *CommandsQueue) PopFront() (cmd Command, err error) {
	return <-q.ch, err
}

//Executor is type intended for control of
//executing of commands
type Executor struct {
	commandsQ *CommandsQueue
}

//PushCommand adds command to executing queue
func (ex *Executor) PushCommand(cmd *Command) {
	ex.commandsQ.PushBack(cmd)
}

//Start of command executing coroutines
func (ex *Executor) Start() {
	ex.commandsQ = NewCommandQueue()
	executing := func() {
		for {
			cmd, err := ex.commandsQ.PopFront()
			if err != nil {
				log.Println("can't pop next command")
			} else {
				cmd.execute()
			}
		}
	}

	go executing()
}
