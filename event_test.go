package event

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	manager := New()
	manager.Add(&Event{
		FN: func(arg ...interface{}) {
			fmt.Println("create cache")
		},
		Name: "create cache",
		Type: "before",
	})

	manager.Add(&Event{
		FN: func(arg ...interface{}) {
			fmt.Println("fetch menu items")
		},
		Name: "fetch menu",
		Type: "before",
	})
	manager.Add(&Event{
		FN: func(arg ...interface{}) {
			fmt.Println("Say Hello")
		},
		Name: "say hello",
		Type: "before",
	})
	manager.Add(&Event{
		FN: func(arg ...interface{}) {
			if arg[0].(string) != "Kemal" {
				t.Error("error")
			}
			fmt.Println("merhaba ", arg[0].(string))
		},

		Args: []interface{}{"Kemal"},
		Name: "mrb",
		Type: "after",
	})

	manager.RunALL(false)

	manager.RunType("after", false)

}
