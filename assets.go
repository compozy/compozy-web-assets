// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "689cd2f9d3685678b7d3adcd52081a89b29511b7ab2da11f137f40f7d317c1d5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a36fa41082a6597522f75beb1936757a1e621ded"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
