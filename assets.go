// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3d68f4acd7b6ea49dc354c6f6a27b286759f9e4604e5a177a204a8b013828db2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6ea5e2612ab0814647d759a1611976f569c23e6f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
