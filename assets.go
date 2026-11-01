// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6413a172eac540d220be8ab593d2b09450f588b869e86457429daae0843c4f1c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "06c329851a7f1e9881f1a3a56b0bd2a66ae8873f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
