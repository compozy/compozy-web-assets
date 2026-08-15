// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a72dfb8f522aeb86dfa9995b6c70554ca913c5611a94ddad716bcfe6f51ab596"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9296ed7e91615b74332931781dd94599ac8db2d5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
