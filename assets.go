// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9f87bdcd97f6a049887699626b083f4bc73678884d377a38e933880fa31e57c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "810f309eb1bc2e756961f3f9955ed934da99112a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
