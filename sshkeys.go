package moulshowcase

import (
	"bufio"
	"fmt"
	"io"

	"github.com/moul/go-sshkeys"
)

func init() {
	RegisterAction("sshkeys", SshkeysAction)
}

func SshkeysAction(qs string, stdin io.Reader) (*ActionResponse, error) {
	scanner := bufio.NewScanner(stdin)
	output := ""
	for scanner.Scan() {
		line := scanner.Text()
		key, err := sshkeys.NewSSHKey([]byte(line))
		if err != nil {
			output += fmt.Sprintf("%v\n", err)
		} else {
			output += key.Hash() + "\n"
		}
	}

	return PlainResponse(output), nil
}
