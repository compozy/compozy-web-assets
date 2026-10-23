// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "962ffd70be467acf714312088059f8a53c3ad256dfe6dc7263ad251a21bfa703"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "922d25ebbb981a0f32f90255b711e2cb080c1fae"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
