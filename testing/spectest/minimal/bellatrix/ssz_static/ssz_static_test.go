package ssz_static

import (
	"testing"

	"github.com/prysmaticlabs/prysm/testing/spectest/shared/bellatrix/ssz_static"
)

func TestMinimal_Merge_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "minimal")
}
