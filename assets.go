// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "afeb658284c7e66fcff4ab0c6340dd0936565795b6f9ab33aaf546d8fe504e4f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fea01695c5cde1a6b87367b9b22ca6f9fbc72ef5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
