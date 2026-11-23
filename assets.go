// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d3ed992aa05102058f53e1b396c043eac88a73158dc8593813bd0866facb4185"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "de2c5e3ed0f9d7aaa95ed61e1f4c4a856c49f2fb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
