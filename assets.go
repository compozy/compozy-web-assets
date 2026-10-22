// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5ee93a5737c1f3e007dd9089531ef27ac389675932674d3c289ca039b76d5336"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "13aeecc4067dc39ba0cb7c0f27aa7b7a64c95954"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
