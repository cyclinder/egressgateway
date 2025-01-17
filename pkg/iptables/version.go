// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package iptables

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major       int
	Minor       int
	Patch       int
	BackendMode string
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return v.Major - other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor - other.Minor
	}
	return v.Patch - other.Patch
}

func ParseVersion(versionString string) (Version, error) {
	re := regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+)`)
	match := re.FindStringSubmatch(versionString)
	if len(match) != 4 {
		return Version{}, fmt.Errorf("invalid version string: %s", versionString)
	}
	major, _ := strconv.Atoi(match[1])
	minor, _ := strconv.Atoi(match[2])
	patch, _ := strconv.Atoi(match[3])
	mode := "legacy"
	if strings.Contains(versionString, "nf_tables") {
		mode = "nft"
	}
	return Version{Major: major, Minor: minor, Patch: patch, BackendMode: mode}, nil
}

func GetVersion() (Version, error) {
	ver := Version{}
	cmd := exec.Command("iptables", "--version")
	out, err := cmd.Output()
	if err != nil {
		return ver, fmt.Errorf("run cmd 'iptables --version' wtith error: %v", err)
	}
	return ParseVersion(string(out))
}
