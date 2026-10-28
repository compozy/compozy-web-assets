// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d02cdecaa7423ce70888de3072424589a712683dd209478eac059bccff895f24"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9c623cbd1a8df175dd9354d0691ad86b0f881b8d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
