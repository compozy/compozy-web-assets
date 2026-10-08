// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2990b6d779db8746cc144bdbb420f68c22dff086395e109d95c6c5fbe5ca3269"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "599886c384b0f6c3b4d6f520982ae205b013d0ca"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
