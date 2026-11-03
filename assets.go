// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e27a914e7c1a66aa46e74b58af03ff5b446ce9aeea55c93ce968b821db2aa291"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ce70672eb781e3db6d9c1a77b7f0ee8daa3d03df"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
