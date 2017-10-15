package inject

import (
	"github.com/neko-neko/godi/internal/container"
)

// DebugLogger publish debugLogger
func (inj *Injector) DebugLogger() Logger {
	return inj.debugLogger
}

// Container publish container
func (inj *Injector) Container() container.Container {
	return inj.container
}
