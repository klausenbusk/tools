package npm

import (
	"context"
	"fmt"
)

// ByTimeStamp implements the sort.Interface for []Version,
// ordering from most recent to least recent time stamps.
type ByTimeStamp []Version

func (a ByTimeStamp) Len() int      { return len(a) }
func (a ByTimeStamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimeStamp) Less(i, j int) bool {
	return a[i].TimeStamp.After(a[j].TimeStamp)
}

// GetMostRecentExistingVersion gets the most recent npm.Version based on time stamp
// that is currently downloaded.
func GetMostRecentExistingVersion(ctx context.Context, existingVersions []string, npmVersions []Version) *Version {
	// create map for fast lookups
	npmMap := make(map[string]Version)
	for _, v := range npmVersions {
		npmMap[v.Version] = v
	}

	// find most recent version
	var mostRecent *Version
	for _, existingVersion := range existingVersions {
		if version, ok := npmMap[existingVersion]; ok {
			if mostRecent == nil || version.TimeStamp.After(mostRecent.TimeStamp) {
				mostRecent = &version // new most recent found
			}
			continue
		}
		// util.Debugf(ctx, "existing version not found on npm: %s", existingVersion)
		panic(fmt.Sprintf("existing version not found on npm: %s", existingVersion))
	}

	return mostRecent
}
