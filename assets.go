// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c97e1af7809cf6d59050f5602a3ead478ef8db70a80cef029fd92ad6e7e67b2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d7ba1a26323e8e5fed3abebf2f737f0068415c63"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
