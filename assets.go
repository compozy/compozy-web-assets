// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "34dbd9784de95d9d9d1aecf988c222ffe0043904a6c24f531bcb80f1638cf566"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5b671cd3ad9620cd779bf9845b4e52aed09966b2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
