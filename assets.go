// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c8b31e5effb4b334c16de3f2fe0da83b6fba0020199248b5356a74f7893b7bd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "22eeb564b4fbecb98e9c8d8977acc036dc3b8984"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
