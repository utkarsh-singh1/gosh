package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/utkarsh-singh1/gosh/inputs"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("% ")

		// read from the keyboard

		input, err := reader.ReadString('\n')

		if err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input

		if err = inputs.ExecInput(input); err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

	}
}
