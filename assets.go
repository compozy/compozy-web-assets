// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f03b4597e0cf0d7315b5df5c88a322ca2bdcb7fafde259e5b5e5b25ee186d94d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9ddda07f1621006a73bd6277b7b5f3bdb2c62078"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
