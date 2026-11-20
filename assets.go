// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "62e0a3d308351bcf8d775a49855beff9e52bff1e8178caac3fe29d9860f3b773"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "070fd958a8bcb79e84c8d53ac11dc5f073a06dbb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
