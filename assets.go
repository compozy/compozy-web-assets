// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b4c8e191dabbf12b9cb5656a261d759b3640a2b1880ab660220cb657cf7d93e6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1533e7b64bb919b8d90a8132efa3f10ad10f3d20"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
