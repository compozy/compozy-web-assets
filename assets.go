// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e215a2f3e6c81f7abeec2aba017add35847a07e6f5e8b08389226e18e11c64da"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6fb2afe080253f9304c7844046dcae969c8a41cb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
