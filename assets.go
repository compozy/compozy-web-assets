// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d9a677a252508f1dc29eb6b4a2316975012a9b6645d8c7f905c19a49cb1eb510"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a8c4a0ae46bcb2475591eac237f82b649172e473"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
