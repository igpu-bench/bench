/*
Copyright © 2023 iGPU Bench Team
*/
package util

import (
	"regexp"

	"github.com/Masterminds/semver/v3"
)

func IsValidVersion(ver string) bool {
	_, err := semver.NewVersion(ver)
	return err == nil
}

// returns true if the two versions are valid and are exactly equal
func IsExactVersionMatch(targetVer string, testVer string) bool {
	targetV, err := semver.NewVersion(targetVer)
	testV, err2 := semver.NewVersion(testVer)

	return testV.Equal(targetV) && err == nil && err2 == nil
}

// returns true if the two versions are valid and are equal within the given precision (tilde comparison)
func IsVersionMatch(targetVer string, testVer string) bool {
	v, _ := semver.NewVersion(testVer)
	c, _ := semver.NewConstraint("~ " + normalizeNoV(targetVer))

	return c.Check(v)
}

func normalizeNoV(ver string) string {
	match := regexp.MustCompile("^v?(.*)$").FindStringSubmatch(ver)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}
