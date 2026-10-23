// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cf96c78edf60a6594b402b7ece018d12d4fa359fdace58a78ae6e0d5a38e1f9f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c68762527a4b29b2210b60e3a2294c78a7affe8f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
