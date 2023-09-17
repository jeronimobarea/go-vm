package main

import (
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	db, err := leveldb.New("", 0, 0, "", false)
	if err != nil {
		panic(err)
	}
	trie.New(nil, db)
	trie.NewDatabase(db, nil)
}
