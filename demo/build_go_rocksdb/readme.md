
# 在window下编译go rocksdb


# ubuntu 

sudo apt install libgflags-dev
sudo apt install libsnappy-dev
sudo apt install zlib1g-dev libbz2-dev liblz4-dev libzstd-dev

sudo apt install librocksdb-dev

CGO_CFLAGS="-I/path/to/rocksdb/include" \
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" \
  go build