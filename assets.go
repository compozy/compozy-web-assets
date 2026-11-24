// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c64478a691fe3e924cce5335faedb8fc9c6cfa7139d2a8ca31fc57ff35782809"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8da2b9572ae2a2b347ba28a1ab14f4235e40635a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
