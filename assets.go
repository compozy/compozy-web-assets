// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e747068f3589d166c1f4d7071207b507c2e3ecc08d77eb91292bf4a303321dde"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8b54e946a4179c6f477c9b24642774df085113da"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
