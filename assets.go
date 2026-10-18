// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4c5f245a495cb5910ff732c60511236d138f4b408f76e1c4b6fcf73b906b4b45"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4ed00e807fa346295377a6b615326a2bb2c916de"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
