package implement

import "strings"

type Cat struct {
}

func (c *Cat) Eat(flag string) string {
	if flag == "true" {
		return strings.ToUpper("milk")
	}
	return "milk"
}

func (c *Cat) Speak(flag string) string {
	if flag == "true" {
		return strings.ToUpper("mee")
	}
	return "mee"
}

func (c *Cat) Move(flag string) string {
	if flag == "true" {
		return strings.ToUpper("run")
	}
	return "run"
}
