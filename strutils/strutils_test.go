package strutils

import (
	"strings"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Log(strings.ToUpper(RandomANI(32)))
}
