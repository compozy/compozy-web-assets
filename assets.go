// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "74564de1cfcb98803b8ceaae03d6e78e6fb0ab843eda0a7183438c341928788b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e39e06bde2bb64d56ac533a1edf4adbf3e59e769"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
