// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "130bb405101cc30911bb176e427187ea585d900fedcddb3a341c3bfac85c6e82"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "93b5d30c1ed52b3f67f0df3edb10c6bc69d9553b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
