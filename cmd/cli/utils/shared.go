package utils

type Blob struct {
	Bytes []byte
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
