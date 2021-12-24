package pet

import (
	"myproject/pet/implement"
	_interface "myproject/pet/interface"
)

const (
	catType  = "cat"
	dogType  = "dog"
	birdType = "bird"
)

var typeMap = make(map[string]interface{})

func init() {
	registry(catType, &implement.Cat{})
	registry(dogType, &implement.Dog{})
	registry(birdType, &implement.Bird{})
}

func registry(name string, form interface{}) {
	if _, ok := typeMap[name]; ok {
		return
	}
	typeMap[name] = form
}

func GetInterface(t string) _interface.Pet {
	form := typeMap[t]
	return form.(_interface.Pet)
}
