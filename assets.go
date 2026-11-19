// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4e6b1effc65d11135e2d363c9a7254c01fdc4e094008564ca28954b364f6bc7d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6ba0445fb8e40b9810c5dad4cfec8a7673e772a4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
