// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "43315335f6ea1be7def3ecb152e6928e3e434b9dda3fedb48127d0248b2b5c65"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "81bb90574930ee39ffeba4a3541f78d8975c274a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
