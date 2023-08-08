package entity

import (
	"fmt"
	"os/exec"

	"github.com/morikuni/failure"
)

type URL string

func (u *URL) Call() error {
	err := exec.Command("open", fmt.Sprintf("%s", *u)).Run()
	if err != nil {
		return failure.Wrap(err)
	}

	return nil
}
