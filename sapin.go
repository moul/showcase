package moulshowcase

import (
	"fmt"

	"github.com/moul/sapin"
)

func init() {
	RegisterAction("sapin", SapinAction)
}

func SapinAction(args []string) (interface{}, error) {
	sapin := sapin.NewSapin(20)
	return fmt.Sprintf("%s", sapin), nil
}
