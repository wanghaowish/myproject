package implement

import "strings"

type Dog struct {
}

func (d *Dog) Eat(flag string) string {
	if flag == "true" {
		return strings.ToUpper("bone")
	}
	return "bone"
}

func (d *Dog) Speak(flag string) string {
	if flag == "true" {
		return strings.ToUpper("woo")
	}
	return "woo"
}

func (d *Dog) Move(flag string) string {
	if flag == "true" {
		return strings.ToUpper("walk")
	}
	return "walk"
}
