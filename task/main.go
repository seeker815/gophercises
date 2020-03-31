package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/seeker815/gophercises/task/cmd"
	"github.com/seeker815/gophercises/task/db"
)

func main() {
	homePath, _ := homedir.Dir()
	dbPath := filepath.Join(homePath, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
