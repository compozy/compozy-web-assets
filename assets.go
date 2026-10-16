// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9729cfc89f53772a9a0633a0ca1b15f679799d8c1e869ae425ef230a129c4b06"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "47abf277c133d1c9f82ebb87fee5556960919407"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
