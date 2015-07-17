package numbers

import (
	"github.com/robertkrimen/otto/registry"
//	"fmt"
)

var entry *registry.Entry = registry.Register(func() string {
	return Source()
})

// Enable numbers runtime inclusion.
func Enable() {
	entry.Enable()
}

// Disable numbers runtime inclusion.
func Disable() {
	entry.Disable()
}

// Source returns the numbers source.
func Source() string {
	res := string(numbers())
//	fmt.Println(res)
	return res
}
