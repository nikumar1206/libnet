package main

import (
	"bufio"
	"fmt"
	"os"
)

// special actions/commands that user can run
var ACTION_CLOSE = "CLOSE\n"     // shutdown the repl and cleanup the underlying connection
var ACTION_CLEAR = "CLEAR\n"     // clear the history
var ACTION_SUSPEND = "SUSPEND\n" // suspend the repl, will keep underlying connection alive.

type REPL struct {
	reader  *bufio.Reader
	prefix  string
	history []string
	conn    Connection
}

func NewREPL(c Connection) *REPL {
	prefix := fmt.Sprintf("%s ❯ ", getFullAddress(c))
	return &REPL{
		conn:    c,
		prefix:  prefix,
		reader:  bufio.NewReader(os.Stdin),
		history: []string{},
	}
}
func (r *REPL) Run() {
	for {
		fmt.Print(r.prefix)
		userInput, _ := r.reader.ReadString('\n')

		switch userInput {
		case ACTION_CLOSE:
			r.Close()
		case ACTION_CLEAR:
			r.Clear()
		case ACTION_SUSPEND:
			panic("Not implemented. Get gud.")
		default:
			// read and write to connection
			r.conn.Write(userInput)
			returnedText := r.conn.Read()

			// add to history and return to user
			r.AddHistory(userInput)
			fmt.Print(returnedText)
		}

	}
}

func (r *REPL) Close() {
	err := r.conn.Close()

	if err != nil {
		handleErr(err)
	}
	os.Exit(0)
}

func (r *REPL) Clear() {
	fmt.Print("\033[H\033[2J")
}

func (r *REPL) AddHistory(userInput string) {
	r.history = append(r.history, userInput)
}
