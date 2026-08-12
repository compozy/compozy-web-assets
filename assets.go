// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5d8edd1927722e1d601d5aa55d6a6e60e82fa11acd988b9b832d57c36586517b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "594d9fdf041e722fca5f60a62351c4c084c71430"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
