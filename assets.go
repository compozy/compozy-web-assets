// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9bfd13fb208d0bcefd1c3aac1f31f51a1350651e3755cc52f988347cc3c127f3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1141d864060678bfec1b75990973fba1b5a87bf1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
