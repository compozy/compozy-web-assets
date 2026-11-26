// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6906c2c9fd3f6512385fd926ee25608a67d28dfc5d34cc0c1572f0c5f4845222"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a5ee302c1978ecfc53794d313c126f9d0a727073"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
