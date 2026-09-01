// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ee33369b364cea923e22acb641a02fe6e84c3bf1ec6d174240f347c03534f52d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b233aead696af20b0432cd7162f6372ccad36b8f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
