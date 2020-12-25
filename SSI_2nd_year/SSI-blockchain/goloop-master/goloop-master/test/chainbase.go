// Code generated by go generate; DO NOT EDIT.
package test

import (
	"context"
	"time"

	"github.com/icon-project/goloop/common/db"
	"github.com/icon-project/goloop/common/log"
	"github.com/icon-project/goloop/module"
)

type ChainBase struct{}

func (_r *ChainBase) Database() db.Database {
	panic("not implemented")
}

func (_r *ChainBase) Wallet() module.Wallet {
	panic("not implemented")
}

func (_r *ChainBase) NID() int {
	panic("not implemented")
}

func (_r *ChainBase) CID() int {
	panic("not implemented")
}

func (_r *ChainBase) NetID() int {
	panic("not implemented")
}

func (_r *ChainBase) Channel() string {
	panic("not implemented")
}

func (_r *ChainBase) ConcurrencyLevel() int {
	panic("not implemented")
}

func (_r *ChainBase) NormalTxPoolSize() int {
	panic("not implemented")
}

func (_r *ChainBase) PatchTxPoolSize() int {
	panic("not implemented")
}

func (_r *ChainBase) MaxBlockTxBytes() int {
	panic("not implemented")
}

func (_r *ChainBase) DefaultWaitTimeout() time.Duration {
	panic("not implemented")
}

func (_r *ChainBase) MaxWaitTimeout() time.Duration {
	panic("not implemented")
}

func (_r *ChainBase) Genesis() []byte {
	panic("not implemented")
}

func (_r *ChainBase) GenesisStorage() module.GenesisStorage {
	panic("not implemented")
}

func (_r *ChainBase) CommitVoteSetDecoder() module.CommitVoteSetDecoder {
	panic("not implemented")
}

func (_r *ChainBase) PatchDecoder() module.PatchDecoder {
	panic("not implemented")
}

func (_r *ChainBase) BlockManager() module.BlockManager {
	panic("not implemented")
}

func (_r *ChainBase) Consensus() module.Consensus {
	panic("not implemented")
}

func (_r *ChainBase) ServiceManager() module.ServiceManager {
	panic("not implemented")
}

func (_r *ChainBase) NetworkManager() module.NetworkManager {
	panic("not implemented")
}

func (_r *ChainBase) Regulator() module.Regulator {
	panic("not implemented")
}

func (_r *ChainBase) Init() error {
	panic("not implemented")
}

func (_r *ChainBase) Start() error {
	panic("not implemented")
}

func (_r *ChainBase) Stop() error {
	panic("not implemented")
}

func (_r *ChainBase) Import(src string, height int64) error {
	panic("not implemented")
}

func (_r *ChainBase) Prune(gs string, dbt string, height int64) error {
	panic("not implemented")
}

func (_r *ChainBase) Backup(file string, extra []string) error {
	panic("not implemented")
}

func (_r *ChainBase) Term() error {
	panic("not implemented")
}

func (_r *ChainBase) State() (string, int64, error) {
	panic("not implemented")
}

func (_r *ChainBase) IsStarted() bool {
	panic("not implemented")
}

func (_r *ChainBase) IsStopped() bool {
	panic("not implemented")
}

func (_r *ChainBase) Reset() error {
	panic("not implemented")
}

func (_r *ChainBase) Verify() error {
	panic("not implemented")
}

func (_r *ChainBase) MetricContext() context.Context {
	panic("not implemented")
}

func (_r *ChainBase) Logger() log.Logger {
	panic("not implemented")
}