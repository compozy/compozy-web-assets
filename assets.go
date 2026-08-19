// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9a824e21173b7e954bdfceca877f732454a01936117be8cff1aea95aa72a17ff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e8df3c29cdc24811670ede99bcd655cbecce4fb9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
