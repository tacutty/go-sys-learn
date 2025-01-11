package main

import "regexp"

func main() {
	r := regexp.MustCompile(`(?i)gopher`)
	tokens := r.Split("I am a Gopher", -1)
	for _, token := range tokens {
		println(token)
	}
}
