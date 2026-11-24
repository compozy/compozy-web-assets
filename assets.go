// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "37ea43d9699e2bfb5c3224616e72659703267d4bf2939542b507b0b92f4353f6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6b64564acafaa8f282954d026359cfaac22304ac"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
