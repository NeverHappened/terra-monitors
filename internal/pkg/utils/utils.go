package utils

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func ValConsToAddr(valcons string) (string, error) {
	_, addr, err := bech32.DecodeAndConvert(valcons)
	return strings.ToUpper(hex.EncodeToString(addr)), err
}

func ValConsPubToAddr(valconspub string) (string, error) {
	_, data, err := bech32.DecodeAndConvert(valconspub)
	if err != nil {
		return "", fmt.Errorf("failed to decode terravalconspub: %w", err)
	}
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKey{},
		ed25519.PubKeyName, nil)
	var pubKey crypto.PubKey
	err = cdc.UnmarshalBinaryBare(data, &pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal binary data to ed25519 pubkey: %w", err)
	}
	return pubKey.Address().String(), nil
}

func GetTerraMonitorsPath() (string, error) {
	dir, err := getCurrentDir()
	if err != nil {
		return "", err
	}

	path, err := getTerraMonitorsPath(dir)
	if err != nil {
		return "", err
	}

	return path, nil
}

func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, err
}

func getTerraMonitorsPath(dir string) (string, error) {
	const dirName = "internal/"

	path := strings.Split(dir, dirName)

	return fmt.Sprintf("%stests/", path[0]), nil
}
