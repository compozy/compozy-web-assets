// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6857fce202c3e06377e0175c80dac3bea3be0b5dfdb95008c633149c3cf1c245"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "de5ab2f4d5703f933f3134d2cf6844d99eb39b9f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
