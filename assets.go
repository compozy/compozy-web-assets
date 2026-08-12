// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e5d91265a7e45cfb90d20929f660d98a23649213ee7333e5bf79fc91280f1400"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "714b7347ba8922eb123ea0c57630cf2a01b92b37"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
