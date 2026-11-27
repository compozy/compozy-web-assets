// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b5f729ec2788ad7f8fd8bc83fcadeec5eaf7a8ee7532b032adec4d2347a8dab1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1f646c482cb7ffc9abf12b772f347bd8fcaca4b7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
