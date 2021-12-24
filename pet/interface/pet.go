package _interface

type Pet interface {
	Eat(flag string) string
	Speak(flag string) string
	Move(flag string) string
}
