package gograveyard

import (
	"regexp"
	"strings"
)

// VerifySemanticVersion based on the recommended regex: https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
func VerifySemanticVersion(ver string) bool {
	ver = strings.TrimPrefix(ver, "v")
	match, _ := regexp.MatchString(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`, ver)
	return match
}
