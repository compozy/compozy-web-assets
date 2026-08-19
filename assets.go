// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2104508242fc8d2dd4bf32667b62152faf075881e6a4b1dbe3cd713fd308fb7b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "502ae6bab6480740f99425569e0b839abde3fe2d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
