// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6f0e0acbc517a26fe656c9328a67993c7dd14a2e2a0671807f228a229bc5000f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "29b4d965e2bb602bc49033bd40ed5f26b31e7012"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
