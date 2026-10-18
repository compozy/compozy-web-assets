// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f06aa80e3c2b01888b223d0840aa597fbd0f4e9e80f72047547c48d417121b1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a32d21bc360d3dd2d5e69093e181e542f2cc0234"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
