// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f49b5121f3a55258dd14271f14c356f02e68c702749a05ff509172599ac1a6dd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cd2824f469caa9a7a8aba8fb5e4ce9518a6fd51d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
