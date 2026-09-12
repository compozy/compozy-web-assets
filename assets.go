// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "817f0fede79dfa15c64eaf8e3f0871f992a92587ed6003e60732f9b0d0bf5f33"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9892912ae2bc895133042f47422f1ec7d0a9f0fa"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
