package implement

import "strings"

type Bird struct {
}

func (b *Bird) Eat(flag string) string {
	if flag == "true" {
		return strings.ToUpper("water")
	}
	return "water"
}

func (b *Bird) Speak(flag string) string {
	if flag == "true" {
		return strings.ToUpper("ling")
	}
	return "ling"
}

func (b *Bird) Move(flag string) string {
	if flag == "true" {
		return strings.ToUpper("fly")
	}
	return "fly"
}
