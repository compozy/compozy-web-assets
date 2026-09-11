// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6ac0fa5b8ea00f08465d277b103eca2de91cbd46d0dcd8ee9af4b9896b53f099"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9a792f51a72009645eb063c6a33a4eb0efc66b3a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
