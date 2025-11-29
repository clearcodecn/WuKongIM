package idgen

import (
	"testing"
	"time"
)

func TestID(t *testing.T) {
	nid := NewTimestampAutoID(1)
	for {
		nid.Next()
		//fmt.Println(nid.Next())
		time.Sleep(200 * time.Millisecond)
	}
}
