// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "97eaf95967b247f753e00997790bdb38427ea0a03e70a22ed3bd67fda2670a71"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3be1703c6c8ed46655fd7f357ca214001bb99e24"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
