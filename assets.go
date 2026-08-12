// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5d8edd1927722e1d601d5aa55d6a6e60e82fa11acd988b9b832d57c36586517b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8947b9b32e58a05f4e4cd3f0b2e24815ee53ce51"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
