package hierarkey

import (
	"errors"
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
	rootNode string
	prevLeaf string
	currLeaf string
	seqMap   map[string]int
}

// NewHierarKey creates a new HierarKey instance
func NewHierarKey(seed int, width int, padding string) *HierarKey {
	if seed <= 0 {
		seed = 1
	}
	if width <= 0 {
		width = 3
	}
	if padding == "" {
		padding = "0"
	}

	h := &HierarKey{
		seed:     seed,
		width:    width,
		padding:  padding,
		rootNode: "ROOT",
		prevLeaf: "",
		seqMap:   make(map[string]int),
	}

	h.currLeaf = h.pad(seed, h.width, h.padding)
	h.seqMap[h.rootNode] = seed - 1

	return h
}

// Validate checks if the hierarKey path string is valid
func (h *HierarKey) Validate(funk string, path string) error {
	matched, err := regexp.MatchString("^[0-9\\.]+$", path)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("%s: %s must contain only Digits and Dots", funk, path)
	}
	if strings.HasPrefix(path, ".") || strings.HasSuffix(path, ".") {
		return fmt.Errorf("%s: %s must NOT start or end with Dots", funk, path)
	}

	parts := strings.Split(path, ".")
	for _, item := range parts {
		if len(item) != h.width {
			return fmt.Errorf("%s: %s of length %d is <> initiated width: %d", funk, item, len(item), h.width)
		}
	}
	return nil
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
func (h *HierarKey) setSeq(path string) error {
	if err := h.Validate("setSeq", path); err != nil {
		return err
	}

	var pathRoot string
	var pathIdx int

	if strings.Index(path, ".") < 0 {
		pathRoot = h.rootNode
		pathIdx, _ = strconv.Atoi(path)
	} else {
		lastDotIdx := strings.LastIndex(path, ".")
		pathRoot = path[:lastDotIdx]
		pathIdx, _ = strconv.Atoi(path[lastDotIdx+1:])
	}

	currIdx, exists := h.seqMap[pathRoot]
	if !exists || pathIdx > currIdx {
		h.seqMap[pathRoot] = pathIdx
	}
	return nil
}

// GetNextSeq gets the next sequence number for a given hierarKey path
func (h *HierarKey) GetNextSeq(path string) (int, error) {
	if err := h.Validate("getNextSeq", path); err != nil {
		return 0, err
	}

	var pathRoot string
	var pathIdx int

	if strings.Index(path, ".") > -1 {
		lastDotIdx := strings.LastIndex(path, ".")
		pathRoot = path[:lastDotIdx]
		pathIdx, _ = strconv.Atoi(path[lastDotIdx+1:])
	} else {
		pathRoot = h.rootNode
		pathIdx, _ = strconv.Atoi(path)
	}

	if currIdx, exists := h.seqMap[pathRoot]; exists {
		nextIdx := currIdx
		if currIdx < pathIdx {
			nextIdx = pathIdx
		} else {
			nextIdx = currIdx + 1
		}
		return nextIdx, nil
	}

	return pathIdx, nil
}

// GetNextLevelSeq gets the next level sequence number for the provided path
func (h *HierarKey) GetNextLevelSeq(path string) (int, error) {
	if err := h.Validate("getNextLevelSeq", path); err != nil {
		return 0, err
	}

	if val, exists := h.seqMap[path]; exists {
		return val + 1, nil
	}
	return h.seed, nil
}

// GetCurrLeaf returns the current leaf
func (h *HierarKey) GetCurrLeaf() string {
	return h.currLeaf
}

// SetCurrLeaf sets the current leaf
func (h *HierarKey) SetCurrLeaf(value string) error {
	if strings.Index(value, ".") > 0 {
		leafArr := strings.Split(value, ".")
		for idx := range leafArr {
			path := strings.Join(leafArr[:idx+1], ".")
			if err := h.setSeq(path); err != nil {
				return err
			}
		}
	} else {
		if err := h.setSeq(value); err != nil {
			return err
		}
	}

	h.prevLeaf = h.currLeaf
	h.currLeaf = value
	return nil
}

// PrevLeaf returns the previous leaf
func (h *HierarKey) PrevLeaf() string {
	return h.prevLeaf
}

// NextLeaf sets the next leaf of hierarKey for a given hierarKey branch
func (h *HierarKey) NextLeaf(currLeaf ...string) (string, error) {
	path := h.currLeaf
	if len(currLeaf) > 0 {
		path = currLeaf[0]
	}

	if err := h.Validate("nextLeaf", path); err != nil {
		return "", err
	}

	idx, err := h.GetNextSeq(path)
	if err != nil {
		return "", err
	}

	if strings.Index(path, ".") < 0 {
		if err := h.SetCurrLeaf(h.pad(idx, h.width, h.padding)); err != nil {
			return "", err
		}
	} else {
		pathRoot := path[:strings.LastIndex(path, ".")]
		if err := h.SetCurrLeaf(fmt.Sprintf("%s.%s", pathRoot, h.pad(idx, h.width, h.padding))); err != nil {
			return "", err
		}
	}

	return h.currLeaf, nil
}

// NextLevel goes to the next level of hierarKey structure
func (h *HierarKey) NextLevel(currLeaf ...string) (string, error) {
	path := h.currLeaf
	if len(currLeaf) > 0 {
		path = currLeaf[0]
	}

	idx, err := h.GetNextLevelSeq(path)
	if err != nil {
		return "", err
	}

	newLeaf := fmt.Sprintf("%s.%s", path, h.pad(idx, h.width, h.padding))
	if err := h.SetCurrLeaf(newLeaf); err != nil {
		return "", err
	}

	return h.currLeaf, nil
}

// PrevLevel goes to the previous level of hierarKey
func (h *HierarKey) PrevLevel(levelDecr ...int) (string, error) {
	decr := 1
	if len(levelDecr) > 0 && levelDecr[0] > 0 {
		decr = levelDecr[0]
	} else if len(levelDecr) > 0 && levelDecr[0] < 1 {
		return "", errors.New(fmt.Sprintf("prevLevel: %d is less than one", levelDecr[0]))
	}

	decrRoot := h.currLeaf
	for i := decr - 1; i >= 0 && strings.Index(decrRoot, ".") > 0; i-- {
		decrRoot = decrRoot[:strings.LastIndex(decrRoot, ".")]
	}

	return h.NextLeaf(decrRoot)
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
func (h *HierarKey) JumpToLevel(path ...string) (string, error) {
	var pathToUse string
	if len(path) > 0 && path[0] != "" {
		pathToUse = h.PadPath(path[0])
	} else {
		pathToUse = h.pad(0, h.width, h.padding)
	}

	if err := h.Validate("jumpToLevel", pathToUse); err != nil {
		return "", err
	}

	return h.NextLeaf(pathToUse)
}
