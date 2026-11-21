// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "688260ddf8acca0128b472e5c5d93542fd0957b8292c3f9fd00e224a993ac270"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2b72438123d53ff772b67d9b8b63d283d6845270"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
