// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bb056723aed47f5b30c1ac9605d56fa02bdcf068cb8d29748774b337aee9e95f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a7eadf4858927494a9fa9a8fae57f164b418c050"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
