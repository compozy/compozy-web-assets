// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a093d8dd34b45b9398ca96012e2c491f7714623720384124700bf5c5685e73c0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "64d4f5afa895e2a62d2cd55a3050d09bbfe600f0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
