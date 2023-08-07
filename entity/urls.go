package entity

import (
	"fmt"
	"os/exec"
)

type URL string

func (u URL) Call() error {
	err := exec.Command("open", fmt.Sprint(u)).Run()
	if err != nil {
		return err
	}

	return nil
}

type URLs []URL