// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2af8cea20a28de08f216b1b627b5d6904a3f4d73523b1695b7f0a74e6d5c4f2b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4140cd2e4144a21dfa1143508f87d9a18b4f55b2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
