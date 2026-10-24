// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4f97d21e32fdaef625a8c9fed8f6405c141432928d94bbce4a5d4043492decc3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5e1e49e49e1b048b2b76ce5f1631d14bacdde317"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
