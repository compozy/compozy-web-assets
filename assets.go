// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f9b77dea4d0c3aa406bcd5a33f3db1f889f3309dabced11fde577a59835b4c6b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "856aa2b7b1a78e8a7f2e78f90a00f1030ad2942d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
