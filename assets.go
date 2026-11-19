// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4e6b1effc65d11135e2d363c9a7254c01fdc4e094008564ca28954b364f6bc7d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "09565e71d8d1abf0a9897fd116dca331719e740b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
