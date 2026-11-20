// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "45b2d5ff4fbc0f729759ff9cb2f71ced933006651e022a1790fc5e687eebf4a6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "16c26e76b3fb0e8644c0942ab6b50d842cd87841"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
