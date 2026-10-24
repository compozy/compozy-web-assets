// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e4996eef4cd621c529b22aa36a2f5f486e615b7d436980a6c736843638b6be84"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a32b58fefd3df53809bca0316c6b0f7b4df0254f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
