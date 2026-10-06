// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a3c9fc66b96a205f7616e498c9a36fc857aa5aaa6a8ba64103133bafdf645ddc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9d842ad2da4cb4ab2980102ddffab96b55a095ce"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
