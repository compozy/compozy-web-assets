// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ba1b58c25b583781164e32e3c1dbe465a87aa8b23a5003264fd5a27d39caaf07"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2faedf5094c807f6bae2f99b2a9e6820872acda4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
