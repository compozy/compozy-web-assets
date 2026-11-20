// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4e6b1effc65d11135e2d363c9a7254c01fdc4e094008564ca28954b364f6bc7d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "94e25adf5fff6882eaf7d7e105f039ba14a1e1e2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
