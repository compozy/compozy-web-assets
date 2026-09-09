// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a18d7588040074211567fcc0de0f601cf95b61f0e2c6b771a8842a50582a037"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cc84920d24b2a6bfbc88882a63cc8c64dfae85a3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
