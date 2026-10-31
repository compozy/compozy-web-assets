// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "719e421e212eb8545dc83745e97f5693fc121ff64b0a8c5b5c2719393de63a07"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "37439a2c9278b2acd0a7e2d46bb3952c3ce880c1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
