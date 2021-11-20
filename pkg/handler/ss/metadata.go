package ss

import (
	"time"

	"github.com/go-gost/gost/pkg/common/util/ss"
	md "github.com/go-gost/gost/pkg/metadata"
	"github.com/shadowsocks/go-shadowsocks2/core"
)

type metadata struct {
	cipher      core.Cipher
	readTimeout time.Duration
	retryCount  int
}

func (h *ssHandler) parseMetadata(md md.Metadata) (err error) {
	const (
		method      = "method"
		password    = "password"
		key         = "key"
		readTimeout = "readTimeout"
		retryCount  = "retry"
	)

	h.md.cipher, err = ss.ShadowCipher(
		md.GetString(method),
		md.GetString(password),
		md.GetString(key),
	)
	if err != nil {
		return
	}

	h.md.readTimeout = md.GetDuration(readTimeout)
	h.md.retryCount = md.GetInt(retryCount)

	return
}