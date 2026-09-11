// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6ac0fa5b8ea00f08465d277b103eca2de91cbd46d0dcd8ee9af4b9896b53f099"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f94f1fc68233627e32274a3adbeb8d956bd4e0b6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
