// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c519a19dadbca8be243b0f268f3498d04a4d864b9c4610a05298373b3c3201c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c6fcb7b7844f7574326b6d860f5c042c153083b1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
