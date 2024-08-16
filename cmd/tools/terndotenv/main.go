package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

/*
* --WRAPPER FOR TERN--
* TERN does not load .env variables
*
* This tool is used to load .env variables through joho/godotenv
* and execute tern migrate command
 */
func main() {
	// load .env variables set in the current folder
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// migrate command
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// run migrate command
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
