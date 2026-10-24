// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3ee5d7ce24b85f0b0b5d4d3ac4c15d42165fec3905676f3bb363676aa76db92a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8660e8bca2f0432b4a8144dd3684ddd59f656ce2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
