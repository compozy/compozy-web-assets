// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0f76874aed72ea2963b6f5e1069ca5079a459485f3da4e6a59e0c713a5bb6226"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2eebafeed98cd7f8bfcf4b766d575ea3d99a2588"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
