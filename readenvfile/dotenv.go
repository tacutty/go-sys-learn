package main

import (
	"fmt"
	"os"
	"flag"
	"os/exec"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	filename := flag.String("filename", ".env", "specify the filename to read")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Error: No command provided to execute. Usage: go run main.go --filename=<env file> <command> [args...]")
	}

	cmdName := flag.Arg(0)
	args := flag.Args()[1:]
	flag.Args()

	cmd := exec.Command(cmdName, args...)

	envs := os.Environ()
	dotenvs, err := godotenv.Read(*filename)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	for k, v := range dotenvs {
		envs = append(envs, k+"="+v)
	}
	cmd.Env = envs
	o, err := cmd.CombinedOutput()
	fmt.Println(string(o))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}