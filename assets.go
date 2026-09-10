// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8a61db40f2f68f09b850161abae70ea78a2379e74b904f24e4f0e4c9bd49a970"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2553622cebb2a2accce3c02d5c821cbfc3ca1d90"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
