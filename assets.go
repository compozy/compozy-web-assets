// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "792990d54d9f9a3ee7b9c653e5feda54e8f0291bd10ab39552f4f7f84f59219f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6e495a69d6e06b049bc4a92f3148be5f43be4dd7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
