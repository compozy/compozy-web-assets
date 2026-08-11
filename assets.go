// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ba1b58c25b583781164e32e3c1dbe465a87aa8b23a5003264fd5a27d39caaf07"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8a1d779bf5af3b22539546622fcc0e586d6bd53d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
