// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "17c36a12ed753eaab58bf3214dc7fe4c30043fdc274e8bf38322b680886f7b0e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7d1f9fd6c900cbdb5e1f5c93b0aa5c2f81ab046a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
