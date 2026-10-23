// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cf96c78edf60a6594b402b7ece018d12d4fa359fdace58a78ae6e0d5a38e1f9f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "271a49aaca902a2d44bb32a8c5ea60647a9b258b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
