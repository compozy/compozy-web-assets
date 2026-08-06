// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "fa7ab1b9e48c6be36086ddebe693785826404ef3ba6ab267af7ab43c8c16e9f1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "62d7e9749bcb7e75e306a72411e7a964cf0a612a"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
