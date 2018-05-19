package command

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

func dbPath() string {
	home, err := homedir.Dir()
	if err != nil {
	}
	return home + "/gotodo.db"
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
