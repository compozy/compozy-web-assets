// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "279bca0c3ea42332865d0b87f4bd81cd0833e247b1ddb5a677eab583bbc3e1d8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1c81d30d080e2c822a25059452bad662651216b5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
