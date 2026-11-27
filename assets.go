// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "64ea8ebd87f9e6ed8c6c09225cece1d0825081e0d904c4db7647f7546ed3c605"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0544c193171d65b7f71a12f554c26cb5289c444d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
