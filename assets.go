// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4e6b1effc65d11135e2d363c9a7254c01fdc4e094008564ca28954b364f6bc7d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2c10bf2f62ae97019ecf2c194346744283c551b2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
