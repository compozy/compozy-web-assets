// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8aa1ae0920160de01f0e8652928da7d363616f8fa40a4e596711011278becfa1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0f55cc51535da12c13abdd8abe45a060050e22b6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
