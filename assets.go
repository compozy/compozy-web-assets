// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ef090bf9a493d6aa0d9576fcd724b470a7700a2199cb72756c88ef60a6834b83"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f1b85557588614dd6e0a16b6358cae12f4252848"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
