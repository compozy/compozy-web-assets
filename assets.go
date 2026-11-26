// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6906c2c9fd3f6512385fd926ee25608a67d28dfc5d34cc0c1572f0c5f4845222"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1a39750cd49007b1783e3ad9ed215a6c7fffcd2b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
