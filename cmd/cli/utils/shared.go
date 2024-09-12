package utils

// TODO: add filepath (think whether it is needed?)
type Blob struct {
	Bytes []byte
}

type Index struct {
	FileToHash map[string]string
}

type Commit struct {
	Message      string
	FileTree     string // hash of the file tree
	ParentCommit string // hash to the previous hash
}

const (
	AddType    = iota // addType = 0
	CommitType = iota // commitType = 1
)

const (
	RootDirName    = ".vgo"
	ObjectsDirName = ".vgo/objects" // serialized blobs
	IndexDirName   = ".vgo/index"   // For staging
    HEADFileName   = ".vgo/HEAD"    // points to the current commit
)
