package moulshowcase

import (
	"fmt"

	"github.com/moul/sapin"
)

func init() {
	RegisterAction("sapin", SapinAction)
}

func SapinAction(args []string) (*ActionResponse, error) {
	sapin := sapin.NewSapin(5)
	return PlainResponse(fmt.Sprintf("%s", sapin)), nil
}
