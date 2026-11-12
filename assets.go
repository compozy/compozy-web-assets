// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "75d954711e92886712d0021821b29eb68a4e72a8bed46c8f7127db6b96787868"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e97bf5922c0c573c72b5b49f890062b06ec01029"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
