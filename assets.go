// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c4bfc3f70c35cd6140460a2d01749713b8df9993f02dc7b3ad4570a4342d6af4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "469359ef95398fcb0f20f32f06163689b6a9c4be"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
