// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e81b9c09cd98678df99279b66670bc8d5379567b472a04b70a04aadbc4acb2da"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "68d03d37f51ef800cff2da2ed56235ebe634dc2f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
