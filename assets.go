// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "67ace54e4bcbc47a2bbaa2df0a49c099c041a851f5af4081b40dc5627ae4d5d5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "610abf2938e4445a012d56a74ffbe8722d62c998"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
