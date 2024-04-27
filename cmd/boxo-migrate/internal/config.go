package migrate

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	ImportPaths map[string]string
	Modules     []string
}

var DefaultConfig = Config{
	ImportPaths: map[string]string{
		"github.com/ipfs/go-bitswap":                     "github.com/littlespeechless/boxo/bitswap",
		"github.com/ipfs/go-ipfs-files":                  "github.com/littlespeechless/boxo/files",
		"github.com/ipfs/tar-utils":                      "github.com/littlespeechless/boxo/tar",
		"github.com/ipfs/interface-go-ipfs-core":         "github.com/littlespeechless/boxo/coreiface",
		"github.com/ipfs/go-unixfs":                      "github.com/littlespeechless/boxo/ipld/unixfs",
		"github.com/ipfs/go-pinning-service-http-client": "github.com/littlespeechless/boxo/pinning/remote/client",
		"github.com/ipfs/go-path":                        "github.com/littlespeechless/boxo/path",
		"github.com/ipfs/go-namesys":                     "github.com/littlespeechless/boxo/namesys",
		"github.com/ipfs/go-mfs":                         "github.com/littlespeechless/boxo/mfs",
		"github.com/ipfs/go-ipfs-provider":               "github.com/littlespeechless/boxo/provider",
		"github.com/ipfs/go-ipfs-pinner":                 "github.com/littlespeechless/boxo/pinning/pinner",
		"github.com/ipfs/go-ipfs-keystore":               "github.com/littlespeechless/boxo/keystore",
		"github.com/ipfs/go-filestore":                   "github.com/littlespeechless/boxo/filestore",
		"github.com/ipfs/go-ipns":                        "github.com/littlespeechless/boxo/ipns",
		"github.com/ipfs/go-blockservice":                "github.com/littlespeechless/boxo/blockservice",
		"github.com/ipfs/go-ipfs-chunker":                "github.com/littlespeechless/boxo/chunker",
		"github.com/ipfs/go-fetcher":                     "github.com/littlespeechless/boxo/fetcher",
		"github.com/ipfs/go-ipfs-blockstore":             "github.com/littlespeechless/boxo/blockstore",
		"github.com/ipfs/go-ipfs-posinfo":                "github.com/littlespeechless/boxo/filestore/posinfo",
		"github.com/ipfs/go-ipfs-util":                   "github.com/littlespeechless/boxo/util",
		"github.com/ipfs/go-ipfs-ds-help":                "github.com/littlespeechless/boxo/datastore/dshelp",
		"github.com/ipfs/go-verifcid":                    "github.com/littlespeechless/boxo/verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline":       "github.com/littlespeechless/boxo/exchange/offline",
		"github.com/ipfs/go-ipfs-routing":                "github.com/littlespeechless/boxo/routing",
		"github.com/ipfs/go-ipfs-exchange-interface":     "github.com/littlespeechless/boxo/exchange",
		"github.com/ipfs/go-merkledag":                   "github.com/littlespeechless/boxo/ipld/merkledag",
		"github.com/boxo/ipld/car":                       "github.com/ipld/go-car",

		// Pre Boxo rename
		"github.com/ipfs/go-libipfs/gateway":               "github.com/littlespeechless/boxo/gateway",
		"github.com/ipfs/go-libipfs/bitswap":               "github.com/littlespeechless/boxo/bitswap",
		"github.com/ipfs/go-libipfs/files":                 "github.com/littlespeechless/boxo/files",
		"github.com/ipfs/go-libipfs/tar":                   "github.com/littlespeechless/boxo/tar",
		"github.com/ipfs/go-libipfs/coreiface":             "github.com/littlespeechless/boxo/coreiface",
		"github.com/ipfs/go-libipfs/unixfs":                "github.com/littlespeechless/boxo/ipld/unixfs",
		"github.com/ipfs/go-libipfs/pinning/remote/client": "github.com/littlespeechless/boxo/pinning/remote/client",
		"github.com/ipfs/go-libipfs/path":                  "github.com/littlespeechless/boxo/path",
		"github.com/ipfs/go-libipfs/namesys":               "github.com/littlespeechless/boxo/namesys",
		"github.com/ipfs/go-libipfs/mfs":                   "github.com/littlespeechless/boxo/mfs",
		"github.com/ipfs/go-libipfs/provider":              "github.com/littlespeechless/boxo/provider",
		"github.com/ipfs/go-libipfs/pinning/pinner":        "github.com/littlespeechless/boxo/pinning/pinner",
		"github.com/ipfs/go-libipfs/keystore":              "github.com/littlespeechless/boxo/keystore",
		"github.com/ipfs/go-libipfs/filestore":             "github.com/littlespeechless/boxo/filestore",
		"github.com/ipfs/go-libipfs/ipns":                  "github.com/littlespeechless/boxo/ipns",
		"github.com/ipfs/go-libipfs/blockservice":          "github.com/littlespeechless/boxo/blockservice",
		"github.com/ipfs/go-libipfs/chunker":               "github.com/littlespeechless/boxo/chunker",
		"github.com/ipfs/go-libipfs/fetcher":               "github.com/littlespeechless/boxo/fetcher",
		"github.com/ipfs/go-libipfs/blockstore":            "github.com/littlespeechless/boxo/blockstore",
		"github.com/ipfs/go-libipfs/filestore/posinfo":     "github.com/littlespeechless/boxo/filestore/posinfo",
		"github.com/ipfs/go-libipfs/util":                  "github.com/littlespeechless/boxo/util",
		"github.com/ipfs/go-libipfs/datastore/dshelp":      "github.com/littlespeechless/boxo/datastore/dshelp",
		"github.com/ipfs/go-libipfs/verifcid":              "github.com/littlespeechless/boxo/verifcid",
		"github.com/ipfs/go-libipfs/exchange/offline":      "github.com/littlespeechless/boxo/exchange/offline",
		"github.com/ipfs/go-libipfs/routing":               "github.com/littlespeechless/boxo/routing",
		"github.com/ipfs/go-libipfs/exchange":              "github.com/littlespeechless/boxo/exchange",

		// Unmigrated things
		"github.com/ipfs/go-libipfs/blocks":       "github.com/ipfs/go-block-format",
		"github.com/littlespeechless/boxo/blocks": "github.com/ipfs/go-block-format",
	},
	Modules: []string{
		"github.com/ipfs/go-bitswap",
		"github.com/ipfs/go-ipfs-files",
		"github.com/ipfs/tar-utils",
		"gihtub.com/ipfs/go-block-format",
		"github.com/ipfs/interface-go-ipfs-core",
		"github.com/ipfs/go-unixfs",
		"github.com/ipfs/go-pinning-service-http-client",
		"github.com/ipfs/go-path",
		"github.com/ipfs/go-namesys",
		"github.com/ipfs/go-mfs",
		"github.com/ipfs/go-ipfs-provider",
		"github.com/ipfs/go-ipfs-pinner",
		"github.com/ipfs/go-ipfs-keystore",
		"github.com/ipfs/go-filestore",
		"github.com/ipfs/go-ipns",
		"github.com/ipfs/go-blockservice",
		"github.com/ipfs/go-ipfs-chunker",
		"github.com/ipfs/go-fetcher",
		"github.com/ipfs/go-ipfs-blockstore",
		"github.com/ipfs/go-ipfs-posinfo",
		"github.com/ipfs/go-ipfs-util",
		"github.com/ipfs/go-ipfs-ds-help",
		"github.com/ipfs/go-verifcid",
		"github.com/ipfs/go-ipfs-exchange-offline",
		"github.com/ipfs/go-ipfs-routing",
		"github.com/ipfs/go-ipfs-exchange-interface",
		"github.com/ipfs/go-libipfs",
	},
}

func ReadConfig(r io.Reader) (Config, error) {
	var config Config
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("reading and decoding config: %w", err)
	}
	return config, nil
}
