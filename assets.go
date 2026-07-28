// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e244a39c0c3a1eb651bf8958a55816134cf5b198df8af166f4db5feec105304a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "44eb84451232308fabba30736410120beb911a0d"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
