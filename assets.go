// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5d5f9b1da4a41a0ffe4a97f804fa5f6ceb3e2c2a4fc2cfc7456133b44a7d6239"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f4583eeb693f2b9f603f58fc2cb54482b3e321fc"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
