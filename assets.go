// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "afeb658284c7e66fcff4ab0c6340dd0936565795b6f9ab33aaf546d8fe504e4f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ae10a8dc027db4ff8dd46186c6eadd3cb22a9d9b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
