// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dd01a63c13e9102237e2ea6f6b23ae84e927bbb4552cff27b227f19e85e37cc9"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "29a596970e450f149c61d9d6da713c3c6c9aa4d4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
