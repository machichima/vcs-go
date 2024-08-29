package utils

// TODO: add filepath (think whether it is needed?)
type Blob struct {
	Bytes []byte
}

type Index struct {
    FileToHash map[string]string
}

const (
	AddType    = iota // addType = 0
	CommitType = iota // commitType = 1
)


const (
	RootDirName    = ".vgo"
	ObjectsDirName = ".vgo/objects" // serialized blobs
	IndexDirName   = ".vgo/index"   // For staging
)
