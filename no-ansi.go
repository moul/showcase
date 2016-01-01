package moulshowcase

import (
	"bufio"
	"fmt"
	"io"

	"github.com/moul/no-ansi"
)

func init() {
	RegisterAction("no-ansi", NoansiAction)
}

func NoansiAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	// Create noansi
	scanner := bufio.NewScanner(stdin)
	output := ""
	for scanner.Scan() {
		line := scanner.Text()
		result, err := noansi.NoAnsiString(line)
		if err != nil {
			output += fmt.Sprintf("Error: %v\n", err)
		} else {
			output += result + "\n"
		}
	}

	return PlainResponse(output), nil
}
