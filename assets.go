// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e2d93b6a6fb76ae9ac1a8ecc33d8c6074fbcff888cb260ae0582d8e5fd8d2557"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2b8f77e2e7e88e7636085a30d4ad7bd4f3ad741b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
