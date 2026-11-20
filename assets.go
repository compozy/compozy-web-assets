// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4dc814356ff3deee30800f497916131f9e9551c205e28d5c14291d6d120ec4ac"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f16a172deddff5fdbb0020cf0e10f5ff6e5ddac7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
