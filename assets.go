// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "eb0154e104f4fd67c5f5b3df027a1d5378c5e2d4484bd9df3b0785f1e55b9696"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8687c03957103ebdcd3670eab3fae446e2e83cf1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
