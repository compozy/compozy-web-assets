// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f06aa80e3c2b01888b223d0840aa597fbd0f4e9e80f72047547c48d417121b1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0a93d83b7cde236a8f0f75bd6b8808e16d3c7d80"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
