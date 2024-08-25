package config

import (
	"common/patch"
	"encoding/hex"
	"strconv"
	"strings"

	cry "github.com/zarkones/xena-crypto"
)

func Patch() (err error) {
	hexEncodedTrustedPubKeyPEM := strings.ReplaceAll(patch.TrustedPubKey, patch.Replacement, "")
	trustedPubKeyPEM, err := hex.DecodeString(hexEncodedTrustedPubKeyPEM)
	if err != nil {
		return err
	}
	TrustedPubKey, err = cry.ImportPubKeyPEM(trustedPubKeyPEM)
	if err != nil {
		return err
	}

	GatewayHost = strings.ReplaceAll(patch.GatewayHost, patch.Replacement, "")

	rawPubSubFlag := strings.ReplaceAll(patch.PubSubFlag, patch.Replacement, "")
	PubSubEnabled = true
	if strings.ToLower(rawPubSubFlag) == "false" {
		PubSubEnabled = false
	}

	PersistAtPath = strings.ReplaceAll(patch.PersistAtPath, patch.Replacement, "")

	maxLoopWait, err := strconv.Atoi(strings.ReplaceAll(patch.MaxLoopWait, patch.Replacement, ""))
	if err != nil {
		return err
	}
	MaxLoopWait = maxLoopWait

	minLoopWait, err := strconv.Atoi(strings.ReplaceAll(patch.MinLoopWait, patch.Replacement, ""))
	if err != nil {
		return err
	}
	MinLoopWait = minLoopWait

	return nil
}
