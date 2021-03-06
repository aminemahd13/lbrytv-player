package version

import "fmt"

var (
	name    = "lbrytv-player"
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

// Name returns main application name
func Name() string {
	return name
}

// Version returns current application version
func Version() string {
	return version
}

// FullName returns current app version, commit and build time
func FullName() string {
	return fmt.Sprintf("%v %v, commit %v, built at %v", Name(), Version(), commit, date)
}
