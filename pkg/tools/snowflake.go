package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/bwmarrin/snowflake"
)

//MD5 计算md5值
func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// 雪花算法的节点
var node *snowflake.Node

//snowFlakeInit 生成一个节点
func init()  {
	if node == nil {
		// Create a new Node with a Node number of 1
		node, _ = snowflake.NewNode(1)
	}
}

// SnowflakeId 获取雪花算法生产的全局唯一id
func SnowflakeId() (string, error) {
	// Generate a snowflake ID.
	id := node.Generate()
	return id.String(), nil
}
