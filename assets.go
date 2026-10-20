// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a0ef1959ab04d1fb333108411657610f7ad20eb432a3af56fece997aca7ccd6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "34e2a6a70b2fdd382269f5be3aff51a382b2dbd4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
