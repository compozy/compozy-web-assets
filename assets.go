// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6906c2c9fd3f6512385fd926ee25608a67d28dfc5d34cc0c1572f0c5f4845222"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "55771102b462488c0e29a6d4363824f0dd40d5d2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
