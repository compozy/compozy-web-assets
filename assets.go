// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "fbee8c0f7a46e7c145a98ec3a0eecf668e3e215ffeaa8440659789413f65b453"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4b0070d4e64054012e4d23a6491db8ec91bbe319"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
