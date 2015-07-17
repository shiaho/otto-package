package moment

import (
	"github.com/robertkrimen/otto/registry"
)

var entry *registry.Entry = registry.Register(func() string {
	return Source()
})

// Enable moment runtime inclusion.
func Enable() {
	entry.Enable()
}

// Disable moment runtime inclusion.
func Disable() {
	entry.Disable()
}

// Source returns the moment source.
func Source() string {
	return string(moment())
}
