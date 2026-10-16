// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4b75b1709eada0bef148c788db7e0c2cc282e9813ac950d8eaa65a6467038aa5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1fbcfe39c9023cd007b4d0ea9732f86b5918612f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
