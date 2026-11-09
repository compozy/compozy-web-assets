// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2ec09cb603b42c996cc079b1416fad3872bd866aed3f6787e3b6bfc1a712ab80"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5f48b457f7e4eaabaccfe08d83952be3dd743d18"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
