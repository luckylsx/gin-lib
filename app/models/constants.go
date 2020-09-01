package models

// 常量定义
type Delete uint8

const (
	// NotDeleted NotDeleted = 0
	NotDeleted Delete = iota
	// NotDeleted NotDeleted = 1
	IsDeleted
)
