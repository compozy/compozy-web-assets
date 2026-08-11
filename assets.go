// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f03b4597e0cf0d7315b5df5c88a322ca2bdcb7fafde259e5b5e5b25ee186d94d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d0afe570cb8b40127f75192e58f58fbe8464e2a7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
