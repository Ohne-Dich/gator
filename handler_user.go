package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login expects a single argument")
	}

	s.cfg.SetUser(cmd.Args[0])

	fmt.Println("User has been set")
	return nil
}
