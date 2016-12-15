/*
Copyright 2016 Helios Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

*/

package version

var (
	// Version is the current version of the Kaigara.
	// The version is of the format Major.Minor.Patch
	Version = "v0.1.0"
)

// GetVersion returns the semver string of the version
func GetVersion() string {
	return Version
}
