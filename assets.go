// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bfbbaf6c8c2b1b4fddb404c3fd19a3308e3d5fae979ff717837adebe40fb4546"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c46be43c732a8fdb06a4fff7157809a4a8e2fa55"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
