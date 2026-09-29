// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6c1d56aaa8a465292740b07b2e0d7f3e428076d8719b32b86bf2b096c2707701"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ecb3abc4e257ab10275e4006d271d000f594f98d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
