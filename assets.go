// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "50e1f53e29fe818c090deda03ca94650f7983d8a170bdb176495e24dc9d2424f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "783376b5df4e77f5cb701046ffedbac2bbea487f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
