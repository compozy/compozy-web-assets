// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8aa1ae0920160de01f0e8652928da7d363616f8fa40a4e596711011278becfa1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "418c1a1880ef033e953fcfec03502e0ef796533b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
