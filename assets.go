// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "22342343c54ad2976d5677ae143f4a4406ae914d91dcbfc766bd3cf0e2cd5f63"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6590dbd07ff2c52b17af1d11bb90f7614a89c7de"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
