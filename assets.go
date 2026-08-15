// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e215a2f3e6c81f7abeec2aba017add35847a07e6f5e8b08389226e18e11c64da"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c9965f022bd4460b104b4cea996da995c093abcd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
