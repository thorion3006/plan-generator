package cli

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Takes the user input from the command line and returns the string.
func UserResponseProcessor() (userResponse string) {
	var err error
	userResponse, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	userResponse = strings.TrimSpace(userResponse)
	reader.Reset(os.Stdin)
	return
}
