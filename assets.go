// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ebdbf34897473462beaac4e002c4908af7e073c39904865c39ef024f482093d8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cd12fe07d63fd6cefdb0fd7a1204e909e178b7af"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
