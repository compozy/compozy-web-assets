// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e670996d8cabbcd68a4930daa5b1e53b6f8a8e88d46b7bf8f287707279f43057"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "13c3a7f8b7eb206b53618896e4259bcda4eaeb40"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
