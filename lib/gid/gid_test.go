package gid

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/oklog/ulid/v2"
	"github.com/teris-io/shortid"
)

func TestShortId(t *testing.T) {
	value, err := shortid.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func TestSnowFlake(t *testing.T) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		// 生成连续的 id
		id := node.Generate()
		fmt.Println(int64(id), " ", id.String())
	}
}

func TestUlid(t *testing.T) {
	fmt.Println(ulid.Make())
}
