// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "559b0b87c06690cff23d282b9cb98635692b7e2d95f1497ea6ee964d04111df6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7b64ab14c6b7c769486092cdb2ea9a0991a85376"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
