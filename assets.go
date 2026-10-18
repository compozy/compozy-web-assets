// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "76c9c10ff226c8944c9fab655ea43fdb3b083015a6bde5d5b3553566e3f5fc68"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "637cbf330d9c36c9f4858cf1a91a5868dde0e042"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
