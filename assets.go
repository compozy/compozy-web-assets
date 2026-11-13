// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bb056723aed47f5b30c1ac9605d56fa02bdcf068cb8d29748774b337aee9e95f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cb3131ae1622a0fac8f0ab44ec778b3961fb0bf9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
