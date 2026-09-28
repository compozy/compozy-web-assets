// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ceefa7a7459f2aa79636be7db7d1a221ab055a22f595732a9dda502bde2c3d7e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "67b86a9b9b3d96f1130e7f8002d0987b924ddef6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
