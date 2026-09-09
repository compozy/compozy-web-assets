// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2d1d07617c465e2376c7e10cb00142d6eb2b901c3f20fc1d1dcd515f2920066c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "744a29acdb8a27c6081251848cdcfa1dbf962f37"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
