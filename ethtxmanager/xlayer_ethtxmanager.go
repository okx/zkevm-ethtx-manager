package ethtxmanager

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/0xPolygon/zkevm-ethtx-manager/etherman"
	"github.com/0xPolygon/zkevm-ethtx-manager/log"
)

// We only use this for signing our tx
func NewWithFrom(cfg Config, from common.Address) (*Client, error) {
	etherman, err := etherman.NewClient(cfg.Etherman)
	if err != nil {
		return nil, err
	}

	storage, err := createStorage(cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	client := Client{
		cfg:      cfg,
		etherman: etherman,
		storage:  storage,
		from:     from,
	}

	log.Init(cfg.Log)

	return &client, nil
}
