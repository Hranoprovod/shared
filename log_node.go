package shared

import (
	"time"
)

// LogNode contains log node data
type LogNode struct {
	Time     time.Time
	Elements *Elements
}

// NewLogNode creates new log node
func NewLogNode(time time.Time, elements *Elements) *LogNode {
	return &LogNode{time, elements}
}
