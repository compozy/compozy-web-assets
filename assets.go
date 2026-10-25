// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c3982a222df77e7a9dc32e19eb44d9701b8ab8ca201920ac88ead263148bb52e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "73be2926201c0f6a7228769e1a0fd3889ad915c7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
