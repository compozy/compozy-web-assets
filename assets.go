// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0d417eb2198e3b372e70e6db665b0a1aceafb8fb9488839105106cc84524e56b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e335272c2fd5c7f6b52293db95e6ad1edb2f9f7f"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
