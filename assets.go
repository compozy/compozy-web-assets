// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0ff15cd43becff38e8539feccd33a8004c3fc7c2f994990ca28d8d97235cc329"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "55ff6b776eddd80a2362052dd53ce6834f6c476e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
