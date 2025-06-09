package hierarkey_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/hierarkey/pkg/hierarkey" // Adjust import path as needed
)

func TestPadFunctions(t *testing.T) {
	hk := hierarkey.NewHierarKey(0, 4)

	if got := hk.Pad(0); got != "0000" {
		t.Errorf("Pad(0) = %q; want %q", got, "0000")
	}
	if got := hk.Pad(7); got != "0007" {
		t.Errorf("Pad(7) = %q; want %q", got, "0007")
	}
	if got := hk.Pad(1000); got != "1000" {
		t.Errorf("Pad(1000) = %q; want %q", got, "1000")
	}
	if got := hk.PadPath("0.5.7.114"); got != "0000.0005.0007.0114" {
		t.Errorf("PadPath('0.5.7.114') = %q; want %q", got, "0000.0005.0007.0114")
	}
}

func TestAddLeafsAndLevels(t *testing.T) {
	hk := hierarkey.NewHierarKey(0, 4)

	if hk.GetCurrLeaf() != "0000" {
		t.Errorf("CurrLeaf = %q; want %q", hk.GetCurrLeaf(), "0000")
	}
	hk.NextLeaf()
	if hk.GetCurrLeaf() != "0000" {
		t.Errorf("CurrLeaf after NextLeaf = %q; want %q", hk.GetCurrLeaf(), "0000")
	}
	hk.NextLeaf()
	if hk.GetCurrLeaf() != "0001" {
		t.Errorf("CurrLeaf after NextLeaf = %q; want %q", hk.GetCurrLeaf(), "0001")
	}
	hk.NextLevel()
	if hk.GetCurrLeaf() != "0001.0000" {
		t.Errorf("CurrLeaf after NextLevel = %q; want %q", hk.GetCurrLeaf(), "0001.0000")
	}
	hk.NextLeaf()
	if hk.GetCurrLeaf() != "0001.0001" {
		t.Errorf("CurrLeaf after NextLeaf x2 = %q; want %q", hk.GetCurrLeaf(), "0001.0001")
	}
	hk.NextLeaf()
	if hk.GetCurrLeaf() != "0001.0002" {
		t.Errorf("CurrLeaf after NextLeaf x2 = %q; want %q", hk.GetCurrLeaf(), "0001.0002")
	}
	hk.NextLevel()
	if hk.GetCurrLeaf() != "0001.0002.0000" {
		t.Errorf("CurrLeaf after NextLevel = %q; want %q", hk.GetCurrLeaf(), "0001.0002.0000")
	}
}

func TestContinueAddingLeafsOnPreviousLevel(t *testing.T) {
	hk := hierarkey.NewHierarKey(0, 4)
	hk.NextLeaf()
	hk.NextLeaf()
	hk.NextLevel()
	hk.NextLeaf()
	hk.NextLeaf()
	hk.NextLevel()

	hk.PrevLevel()
	if hk.GetCurrLeaf() != "0001.0003" {
		t.Errorf("CurrLeaf after PrevLevel = %q; want %q", hk.GetCurrLeaf(), "0001.0003")
	}
	hk.PrevLevel()
	if hk.GetCurrLeaf() != "0002" {
		t.Errorf("CurrLeaf after PrevLevel = %q; want %q", hk.GetCurrLeaf(), "0002")
	}
}

func TestJumpAndAddLeafs(t *testing.T) {
	hk1 := hierarkey.NewHierarKey(1, 4)
	leaf := hk1.JumpToLevel("0007.0006")
	if leaf != "0007.0006" {
		t.Errorf("JumpToLevel = %q; want %q", leaf, "0007.0006")
	}
	leaf = hk1.NextLevel()
	if leaf != "0007.0006.0001" {
		t.Errorf("NextLevel = %q; want %q", leaf, "0007.0006.0001")
	}
	if hk1.GetCurrLeaf() != "0007.0006.0001" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0006.0001")
	}
	hk1.JumpToLevel("0001.0004")
	if hk1.GetCurrLeaf() != "0001.0004" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0001.0004")
	}
	hk1.NextLeaf()
	if hk1.GetCurrLeaf() != "0001.0005" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0001.0005")
	}
	hk1.JumpToLevel("0007.0006.0005")
	if hk1.GetCurrLeaf() != "0007.0006.0005" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0006.0005")
	}
	hk1.JumpToLevel("0007.0006.0005")
	if hk1.GetCurrLeaf() != "0007.0006.0006" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0006.0006")
	}
	hk1.NextLevel()
	hk1.JumpToLevel("0007.0006.0005.0001")
	if hk1.GetCurrLeaf() != "0007.0006.0005.0001" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0006.0005.0001")
	}
}

func TestContinueAddingLeafsOnPreviousLevel2(t *testing.T) {
	hk1 := hierarkey.NewHierarKey(1, 4)
	hk1.JumpToLevel("0007.0006.0005.0001")
	hk1.PrevLevel()
	if hk1.GetCurrLeaf() != "0007.0006.0006" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0006.0006")
	}
	hk1.PrevLevel()
	if hk1.GetCurrLeaf() != "0007.0007" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0007.0007")
	}
}

func TestAddAdditionalRootLeafs(t *testing.T) {
	hk1 := hierarkey.NewHierarKey(1, 4)
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.NextLeaf()
	hk1.JumpToLevel("0000")
	if hk1.GetCurrLeaf() != "0008" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0008")
	}
	hk1.NextLeaf()
	if hk1.GetCurrLeaf() != "0009" {
		t.Errorf("CurrLeaf = %q; want %q", hk1.GetCurrLeaf(), "0009")
	}
	hk1.JumpToLevel("")
}

func TestPathPunctuationError(t *testing.T) {
	hk2 := hierarkey.NewHierarKey(1, 4)
	hk2.JumpToLevel(".0007.0006.")
	if hk2.GetCurrLeaf() != "0001.0007.0006.0001" {
		t.Errorf("CurrLeaf = %q; want %q", hk2.GetCurrLeaf(), "0001.0007.0006.0001")
	}
}

/*
func TestDigitCharacterError(t *testing.T) {
	hk2 := hierarkey.NewHierarKey(1, 4)
	res:= hk2.JumpToLevel("0007.0A06")
	if !regexp.MustCompile(`must contain only Digits and Dots`).MatchString(fmt.Sprintf("%v", res)) {
		t.Errorf("Expected digit character error, got %v", res)
	}
}
*/
/*
func TestTooLargeFieldError(t *testing.T) {
	hk2 := hierarkey.NewHierarKey(1, 4)
	res := hk2.JumpToLevel("00087.0006")
	if !regexp.MustCompile(`initiated width`).MatchString(fmt.Sprintf("%v", res)) {
		t.Errorf("Expected too large field error, got %v", res)
	}
}

func TestOutOfBoundsIncrement(t *testing.T) {
	hk2 := hierarkey.NewHierarKey(1, 4)
	hk2.JumpToLevel("0007.0006.9999")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic or error on out of bounds increment")
		}
	}()
	hk2.NextLeaf()
}
*/
func randomIntFromInterval(r *rand.Rand, min, max int) int {
	return r.Intn(max-min+1) + min
}

func TestHierarKeyCanGenerate10000Entries(t *testing.T) {
	hk5 := hierarkey.NewHierarKey(1, 5)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		pathDepth := randomIntFromInterval(r, 1, 21)
		path := ""
		for j := 0; j < pathDepth; j++ {
			num := randomIntFromInterval(r, 0, 9000)
			if j == 0 {
				path = fmt.Sprintf("%d", num)
			} else {
				path = fmt.Sprintf("%s.%d", path, num)
			}
		}
		hk5.JumpToLevel(path)
	}
}
