package version

import "runtime"

var (
	Version       string
	Commit        string
	BuildDatetime string
)

type BuildInfo struct {
	Version       string
	Commit        string
	GoVersion     string
	BuildDatetime string
}

func Info() BuildInfo {
	return BuildInfo{
		Version:       Version,
		Commit:        Commit,
		GoVersion:     runtime.Version(),
		BuildDatetime: BuildDatetime,
	}
}
