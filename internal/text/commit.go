package text

import "encoding/hex"

// Commit is a parsed commit that contains information about category, scope and heading
type Commit struct {
	Category string
	Heading  string
	Body     string
	Scope    string
	Hash     Hash
	Issues   []int
}

// Hash describes a commit hash
type Hash [20]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}
