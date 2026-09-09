// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "97eaf95967b247f753e00997790bdb38427ea0a03e70a22ed3bd67fda2670a71"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9b422f1df344f659ac174c749e8a6ab3b84f110d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
