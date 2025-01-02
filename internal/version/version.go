package version

import "runtime"

var (
	Version string
	Commit  string
)

type BuildInfo struct {
	Version   string
	Commit    string
	GoVersion string
}

func Info() BuildInfo {
	return BuildInfo{
		Version:   Version,
		Commit:    Commit,
		GoVersion: runtime.Version(),
	}
}
