// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8d34fd01b8fc178e4b75412731ae73516195798a4e496ab1b9e245196cac8ad2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "10a375345dea9250b09df2a34b50a6363fa5716b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
