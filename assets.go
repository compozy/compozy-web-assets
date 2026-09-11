// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6ac0fa5b8ea00f08465d277b103eca2de91cbd46d0dcd8ee9af4b9896b53f099"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e45b472b4f25afa90445459a727fb456834725b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
