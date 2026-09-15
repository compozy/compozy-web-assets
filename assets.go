// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7f7c41bdc9865c362d1fcfbeec5d087dafd92a36c583a20672da4cf7e2055251"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "09648e093e2d68751cf2ee4c218f0293c8fde354"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
