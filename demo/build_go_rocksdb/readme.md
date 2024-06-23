
# 在window下编译go rocksdb


# ubuntu 

<!-- sudo apt install libgflags-dev -->
sudo apt install libsnappy-dev
sudo apt install zlib1g-dev libbz2-dev liblz4-dev libzstd-dev

sudo apt install librocksdb-dev

CGO_CFLAGS="-I/usr/include/rocksdb/" \
CGO_LDFLAGS="-L/usr/lib/ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
  go build