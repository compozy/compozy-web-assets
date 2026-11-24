// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "459369b6ca09585e4483a1147cd6ba3dcdece32482cee896657d2d463f811902"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "288e00db7b4e8337db14f7143f7183953bd06ec0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
