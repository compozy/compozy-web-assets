// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "17c36a12ed753eaab58bf3214dc7fe4c30043fdc274e8bf38322b680886f7b0e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f21475050c884b16dfc41eb819b011f11da2785e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
