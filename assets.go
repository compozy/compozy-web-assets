// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2104508242fc8d2dd4bf32667b62152faf075881e6a4b1dbe3cd713fd308fb7b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "afb1b96ea3e99b8a81eefa01b949fe4d88146bdd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
