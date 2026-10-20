// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c7a881b887064b6a89e5d757b87ea250b3a20d0ce13f91c10d0a0c033d76ff40"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f6331360e2c7dd445bc65ec7f225069b861425f9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
