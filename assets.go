// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "97eaf95967b247f753e00997790bdb38427ea0a03e70a22ed3bd67fda2670a71"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cb9f5ce0a0676bf9af326075abeae0b620b7eae8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
