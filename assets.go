// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b3200762f2589945690124cff300637c608d0af39389745e0dbc2c53aeedd3d4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e6457afe0629d012928b3e7880eb600f642696cc"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
