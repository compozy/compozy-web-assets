// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "39e9cd406fa72245ed9a8d5552be8e58a3f091ba930d65d39270d8e85d9755b7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "332d314ac14f4ffdf69238978d7e8ea312d9c42d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
