// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "00a36abc4f9c14ac4a4d9e3738ec9aaf8ec6b4a8fa0e63f323cbc3e5ca753092"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "df4c4274c3e32f81fb1738319f8379f8f065c140"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
