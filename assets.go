// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "13de0ea9b84e5126c6dde27ad11b1338eb1224fa97a5368788a4edebc5eced25"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "524a8c33500ae388d4b5d0dd7840f536a31ec7de"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
