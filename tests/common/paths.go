package common

import "runtime"

var SLASH = func() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}()

const PROJECT_EXPORT_PATH = "_export"
const C2_EXPORT_PATH = PROJECT_EXPORT_PATH + "/c2"
const UI_EXPORT_PATH = PROJECT_EXPORT_PATH
const ROOT_PATH_RELATIVE_TO_TESTER = ".."
const AGENT_EXPORT_PATH = "common/builder/static/agents"
