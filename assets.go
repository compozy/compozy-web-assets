// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a879eb073172358e654a1613ebc2ba81d385e23c0f94f92a761f802924124bd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5ae7dbbfe90c9b9c6ce8f7821d4148d9bb21d882"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
