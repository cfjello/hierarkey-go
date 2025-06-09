package hierarkey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// HierarKey generates a hierarchical tree numbering such as: '0003.0001.0004'
type HierarKey struct {
	seed     int
	width    int
	padding  string
	prevLeaf string
	currLeaf string
	seqMap   map[string]int
}

// NewHierarKey creates a new HierarKey instance
func NewHierarKey(seed int, width int, padding ...string) *HierarKey {
	if seed < 0 {
		seed = 1
	}
	if width <= 0 {
		width = 3
	}

	// Default padding is "0" if not provided
	padChar := "0"
	if len(padding) > 0 && padding[0] != "" {
		padChar = padding[0]
	}

	h := &HierarKey{
		seed:    seed,
		width:   width,
		padding: padChar,
		// rootNode: "ROOT",
		prevLeaf: "",
		seqMap:   make(map[string]int),
	}

	h.currLeaf = h.pad(seed, h.width, h.padding)
	h.seqMap[h.currLeaf] = seed - 1

	return h
}

// Validate checks if the hierarKey path string is valid
func (h *HierarKey) Validate(funk string, path string) {
	matched, err := regexp.MatchString("^[0-9\\.]+$", path)
	if err != nil {
		panic(err)
	}
	if !matched {
		panic(fmt.Sprintf("%s: %s must contain only Digits and Dots", funk, path))
	}
	if strings.HasPrefix(path, ".") || strings.HasSuffix(path, ".") {
		panic(fmt.Sprintf("%s: %s must NOT start or end with Dots", funk, path))
	}

	parts := strings.Split(path, ".")
	for _, item := range parts {
		if len(item) != h.width {
			panic(fmt.Sprintf("%s: %s of length %d is <> initiated width: %d", funk, item, len(item), h.width))
		}
	}
}

// Pad formats a number with leading zeros
func (h *HierarKey) pad(n int, w int, z string) string {
	nStr := strconv.Itoa(n)
	if len(nStr) >= w {
		return nStr
	}
	return strings.Repeat(z, w-len(nStr)) + nStr
}

// SetSeq remembers the hierarKey root and sets the next sequence number
func (h *HierarKey) setSeq(path string, idx int) {
	h.seqMap[path] = idx
}

// GetNextSeq gets the next sequence number for a given hierarKey path
func (h *HierarKey) GetNextSeq(path string) int {
	h.Validate("getNextSeq", path)

	var pathRoot string
	var pathIdx int

	if strings.Index(path, ".") > 0 {
		lastDotIdx := strings.LastIndex(path, ".")
		pathRoot = path[:lastDotIdx]
		pathIdx, _ = strconv.Atoi(path[lastDotIdx+1:])
	} else {
		pathRoot = ""
		pathIdx, _ = strconv.Atoi(path)
	}

	goOn := true
	currPath := path
	nextIdx := h.seed

	for goOn {
		newIdx, exists := h.seqMap[currPath]
		if exists {
			nextIdx = newIdx + 1
			if pathRoot != "" {
				currPath = fmt.Sprintf("%s.%s", pathRoot, h.pad(nextIdx, h.width, h.padding))
			} else {
				currPath = h.pad(nextIdx, h.width, h.padding)
			}
		} else {
			goOn = false
			break
		}
		// Check for a change in nextIdx
		if nextIdx == pathIdx {
			goOn = false
			break
		}
	}
	return nextIdx
}

// GetNextLevelSeq gets the next level sequence number for the provided path
func (h *HierarKey) GetNextLevelSeq(path string) int {
	h.Validate("getNextLevelSeq", path)

	if path == "" {
		path = fmt.Sprintf("%s.%s", h.currLeaf, h.pad(h.seed, h.width, h.padding))
	}
	return h.GetNextSeq(path)
}

// GetCurrLeaf returns the current leaf
func (h *HierarKey) GetCurrLeaf() string {
	return h.currLeaf
}

// SetCurrLeaf sets the current leaf
func (h *HierarKey) SetCurrLeaf(value string, idxIn int) {
	h.prevLeaf = h.currLeaf
	h.currLeaf = value
	h.setSeq(value, idxIn)
}

// PrevLeaf returns the previous leaf
func (h *HierarKey) PrevLeaf() string {
	return h.prevLeaf
}

// NextLeaf sets the next leaf of hierarKey for a given hierarKey branch
func (h *HierarKey) NextLeaf(currLeaf ...string) string {
	path := ""
	if len(currLeaf) > 0 {
		path = currLeaf[0]
	} else {
		path = h.currLeaf
	}
	h.Validate("nextLeaf", path)

	idx := h.GetNextSeq(path)

	if !strings.Contains(path, ".") {
		h.SetCurrLeaf(h.pad(idx, h.width, h.padding), idx)
	} else {
		pathRoot := path[:strings.LastIndex(path, ".")]
		h.SetCurrLeaf(fmt.Sprintf("%s.%s", pathRoot, h.pad(idx, h.width, h.padding)), idx)
	}
	return h.currLeaf
}

// NextLevel goes to the next level of hierarKey structure
func (h *HierarKey) NextLevel(currLeaf ...string) string {
	newLeaf := h.currLeaf
	if len(currLeaf) > 0 {
		newLeaf = currLeaf[0]
	}
	newLeaf = fmt.Sprintf("%s.%s", newLeaf, h.pad(h.seed, h.width, h.padding))
	idx := h.GetNextLevelSeq(newLeaf)
	h.SetCurrLeaf(newLeaf, idx)
	return h.currLeaf
}

// PrevLevel goes to the previous level of hierarKey
func (h *HierarKey) PrevLevel(levelDecr ...int) string {
	decr := 1
	if len(levelDecr) > 0 && levelDecr[0] > 0 {
		decr = levelDecr[0]
	} else if len(levelDecr) > 0 && levelDecr[0] < 1 {
		panic(fmt.Sprintf("prevLevel: %d is less than one", levelDecr[0]))
	}
	decrRoot := h.currLeaf
	for i := decr - 1; i >= 0 && strings.Index(decrRoot, ".") > 0; i-- {
		decrRoot = decrRoot[:strings.LastIndex(decrRoot, ".")]
	}
	return h.NextLeaf(decrRoot)
}

// pad formats a number with leading padding characters to the specified width
func (h *HierarKey) Pad(n int, z ...string) string {
	padChar := h.padding
	if len(z) > 0 {
		padChar = z[0]
	}
	nStr := strconv.Itoa(n)
	return strings.Repeat(padChar, h.width-len(nStr)) + nStr
}

// PadPath provides uniform padding for each dot-separated entry
func (h *HierarKey) PadPath(path string) string {
	parts := strings.Split(path, ".")
	for i, item := range parts {
		parts[i] = strings.Repeat(h.padding, h.width-len(item)) + item
	}
	return strings.Join(parts, ".")
}

// JumpToLevel jumps to a level of the hierarKey structure
func (h *HierarKey) JumpToLevel(path ...string) string {
	var pathToUse string
	if len(path) > 0 && path[0] != "" {
		pathToUse = h.PadPath(path[0])
	} else {
		pathToUse = h.pad(0, h.width, h.padding)
	}
	h.Validate("jumpToLevel", pathToUse)
	// Split the path and check if each parent exists in h.seqMap
	parts := strings.Split(pathToUse, ".")
	parent := ""
	for i := 0; i < len(parts); i++ {
		if i == 0 {
			parent = parts[i]
		} else {
			parent = fmt.Sprintf("%s.%s", parent, parts[i])
		}
		if _, exists := h.seqMap[parent]; !exists {
			// If a parent doesn't exist in the map, add it with initial value
			part := parts[i]
			pathIdx, err := strconv.Atoi(part)
			if err != nil {
				panic(fmt.Sprintf("jumpToLevel: failed to convert %s to integer: %v", part, err))
			}
			h.SetCurrLeaf(parent, pathIdx)
		} else if i == len(parts)-1 {
			// If it's the last part and it already exists, then set it as the current leaf
			// to the next available sequence number on that path level
			return h.NextLeaf(parent)
		}
	}
	return h.currLeaf
}
