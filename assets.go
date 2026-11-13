// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4f0852f8f0069b088371fdedc34baee8f66dcdf261c61402bec3e2ef391c344e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "adbe71f960b2ce49e0ae629aa28e891a18e1e0b3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
