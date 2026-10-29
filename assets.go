// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5a1883ee9895f0f0de1c897583dae77a8b82fd7d1a24fc2bb1ad9a0aa873b0fc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4f73f3c770720930b6463be8a5a3ddefb559b9d5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
