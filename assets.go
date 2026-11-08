// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a59fda0ecbeb8aa79f348a3cc96c540ddb6709c47f5ee2e8c1ff8afbd52df668"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "255177e4c0b06d1e822cd53a2c88303a2b76c540"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
