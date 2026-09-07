// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "97eaf95967b247f753e00997790bdb38427ea0a03e70a22ed3bd67fda2670a71"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "eacd0cb53b9cf4b72fb61fb5db6a75868a3a48ce"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
