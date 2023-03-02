
package main

func main(){
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	db, err := grocksdb.OpenDb(opts, "temp/db")
}