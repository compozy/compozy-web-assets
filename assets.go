// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "64a39be87c63d1e66ce3368c1fbef986e97ded76f565a53277a7f44af0ae10b3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cfd16f1ed535244fc145092c15c5e50a2b8acad7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
