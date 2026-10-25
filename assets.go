// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a09cc561965a48c0128ab836044f7c6e9da7b4e52f99265a8f43ebe96ef2b89e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "136a188395f8b1f85caf44dc30abe2dd674a4af5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
