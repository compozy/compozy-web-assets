// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d06839a5d4bdc002cc0096e24ed441c24cee2015228f9e78beb89a3d434f537a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9221d47f49366ffb952625f6904a799c38f38024"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
