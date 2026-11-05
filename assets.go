// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3ec825f0b9300d53472f77683586948368e8bfb54e4cb743ea8fc79ee8e0fe63"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "708807765c89a568f14d4b04ecf99d67a864852b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
