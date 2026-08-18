// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2104508242fc8d2dd4bf32667b62152faf075881e6a4b1dbe3cd713fd308fb7b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f003f31d6e4695ec36093323edf2ca8c8a27fe8b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
