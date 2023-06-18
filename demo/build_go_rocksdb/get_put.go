package main

import "github.com/linxGnu/grocksdb"

func main() {
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	db, err := grocksdb.OpenDb(opts, "temp/db")
	_ = db
	_ = err
}
