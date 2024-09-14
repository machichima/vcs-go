package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirs(paths ...string) error {
	for _, path := range paths {
		if err := CreateOneDir(path); err != nil {
			return err
		}
	}
	return nil
}

func GetFiles(path string) ([]string, error) {

	var fileNames []string

	// queue
	queue := make([]string, 0)
	queue = append(queue, path)

	for len(queue) > 0 {

		// pop the first element
		path := queue[0]
		queue = queue[1:]

		dirs, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				if dir.Name() != ".vgo" {
					queue = append(queue, path+dir.Name()+"/")
				}
			} else {
				fileNames = append(fileNames, path+dir.Name())
			}
		}
	}

	// fmt.Println(fileNames)

	return fileNames, nil
}

// create directory and all parents
func CreateOneDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func FileToStruct(path string) (Blob, error) {
	// Read the file to be serialized
	readBytes, err := os.ReadFile(path)
	if err != nil {
		return Blob{}, err
	}

	blob := Blob{Bytes: readBytes}

	return blob, nil

}

// Save files with hash string as file name
// and in the folder names the first two characters
// of the hash. The files will be saved in the ObjectsDirName dir (.vgo/objects)
//
// e.g. hash1: dfq8hfroihffjlkasj / hash2: df32fjqf81efh1ofj
// saved in df/dfq8hfroihffjlkasj and df/df32fjqf81efh1ofj file
// If the commandType is AddType, then add the file name and hash to the index file
//
// if error is nil, return true means added the new file, else return false
func SaveFileByHash(filePath string, hash string, blob []byte, commandType int) (bool, error) {

	// create parent dir
	parentDir := hash[:2]
	fullObjectsDir := ObjectsDirName + "/" + parentDir

	if err := CreateOneDir(fullObjectsDir); err != nil {
		return false, err
	}

	// write blob to file
	if err := os.WriteFile(fullObjectsDir+"/"+hash, blob, os.ModePerm); err != nil {
		return false, err
	}

	if commandType == AddType {
		//read index
		index, err := ReadIndexFile()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// if the index file does not exist, create one with empty index struct
				fmt.Println("Index file does not exist, creating one new")
                index = Index{FileToHash: make(map[string]string)}
			} else {
				return false, err
			}
		}
        isNewFile, err := AddToIndex(&index, filePath, hash)
        if err != nil {
            return false, err
        }

        // Write index file if the file is new
        // TODO: update this to write index file for all added files
        if isNewFile {
            if err := WriteIndexFile(index); err != nil {
                return false, err
            }
        }

        return isNewFile, nil
	}

	return false, nil
}

// read INDEX file in .vgo if exist else create one
//
// Return the Index structure and error
func ReadIndexFile() (Index, error) {
	// Deserialize the index file to Index struct

	byte, err := os.ReadFile(IndexDirName)
	if err != nil {
		return Index{}, err
	}

	buff := bytes.NewBuffer(byte)

	// deserailize the byte to Index struct

	index, err := DeserializeIndex(buff)
	if err != nil {
		return Index{}, err
	}

	return index, nil
}

// write the index struct to INDEX file
func WriteIndexFile(index Index) error {

	// serialized index buffer
	var serializedBuffer bytes.Buffer
	if err := SerializeIndex(index, &serializedBuffer); err != nil {
		return err
	}

	// write the buffer to INDEX file
	if err := os.WriteFile(IndexDirName, serializedBuffer.Bytes(), os.ModePerm); err != nil {
		return err
	}

	return nil
}


// Add file and its hash to INDEX file. Serialize the index struct and write to INDEX file
//
// Return bool and error, true if the file is added to index, false if the file already exists in index.
// If err is not nil, bool is false
func AddToIndex(index *Index, file string, hash string) (bool, error) {

    var isNewFile bool = false

	// file already exists in index
	if index.FileToHash[file] != "" {
        if index.FileToHash[file] == hash {
            // file already exists in index with same hash
            return false, nil
        }
		if err := DeleteObject(index.FileToHash[file]); err != nil {
			return false, err
		}
	} else {
        isNewFile = true
    }

    // empty index content (no staged files)
    if len(index.FileToHash) == 0 {
        index.FileToHash = make(map[string]string)
    }

	// update or add file-hash to index
	index.FileToHash[file] = hash

	return isNewFile, nil
}

func DeleteObject(hash string) error {
	path := filepath.Join(ObjectsDirName, hash[:2], hash)
	if err := os.Remove(path); err != nil {
		return err
	}

	f, err := os.Open(filepath.Dir(path))
	if err != nil {
		return err
	}

	if _, err := f.ReadDir(1); err == io.EOF {
		// the foler is empty, delete the folder
		defer os.Remove(filepath.Dir(path))
	}

	return nil
}


// write file tree (index object) to objects, returning the hash of the filetree if error is nil
func WriteFileTree(index Index) (string, error) {

	// serialized index buffer
	var serializedBuffer bytes.Buffer
	if err := SerializeIndex(index, &serializedBuffer); err != nil {
		return "", err
	}

    hash, err := HashBlob(serializedBuffer.Bytes())
    if err != nil {
        return "", err
    }


	parentDir := hash[:2]
	fullObjectsDir := ObjectsDirName + "/" + parentDir

	if err := CreateOneDir(fullObjectsDir); err != nil {
		return "", err
	}

	// write blob to file
	if err := os.WriteFile(fullObjectsDir+"/"+hash, serializedBuffer.Bytes(), os.ModePerm); err != nil {
		return "", err
	}

	return hash, nil
}


// Read the filetree from the provided hash (within folder .vgo/objects/hash[:2])
//
// Return index 
func ReadFileTree(hash string) (Index, error) {

	parentDir := hash[:2]
	fullObjectsDir := ObjectsDirName + "/" + parentDir

    byte, err := os.ReadFile(filepath.Join(fullObjectsDir, hash))
	if err != nil {
		return Index{}, err
	}

	buff := bytes.NewBuffer(byte)
	index, err := DeserializeIndex(buff)
	if err != nil {
		return Index{}, err
	}

	return index, nil
}


// write the commit struct to file and return its hash if err is nil
func WriteCommit(commit Commit) (string, error) {

	// serialized index buffer
	var serializedBuffer bytes.Buffer
	if err := SerializeCommit(commit, &serializedBuffer); err != nil {
		return "", err
	}

    hash, err := HashBlob(serializedBuffer.Bytes())
    if err != nil {
        return "", err
    }

	parentDir := hash[:2]
	fullObjectsDir := ObjectsDirName + "/" + parentDir

	if err := CreateOneDir(fullObjectsDir); err != nil {
		return "", err
	}

	// write blob to file
	if err := os.WriteFile(fullObjectsDir+"/"+hash, serializedBuffer.Bytes(), os.ModePerm); err != nil {
		return "", err
	}

	return hash, nil
}


// Read the commit from the provided hash (within folder .vgo/objects/hash[:2])
//
// Return Commit struct
func ReadCommit(hash string) (Commit, error) {

    fullObjectsDir := filepath.Join(ObjectsDirName, hash[:2])

    byte, err := os.ReadFile(filepath.Join(fullObjectsDir, hash))
	if err != nil {
		return Commit{}, err
	}

	buff := bytes.NewBuffer(byte)
    commit, err := DeserializeCommit(buff)
	if err != nil {
		return Commit{}, err
	}

	return commit, nil
}

// write the blob by serialize the file and save under
// the file with name of the file's hash
//
// The file and their hash will be written into the Index file
// for recording staged files
//
// Return bool showing whether the added file is new 
// (not already in staging). If error is nil
// true is added file is new, else false
func WriteFileBlobWithSerialize(filePath string) (bool, error) {
    var blob Blob
    var err error
    blob.Bytes, err = os.ReadFile(filePath)
    if err != nil {
        return false, err
    }

	var serBlob bytes.Buffer

	if err := SerializeBlob(blob, &serBlob); err != nil {
		return false, err
	}

	hash, err := HashBlob(serBlob.Bytes())
	if err != nil {
		return false, err
	}

    isNewFile, err := SaveFileByHash(filePath, hash, serBlob.Bytes())
    if err != nil {
        return false, err
    }

    return isNewFile, nil
}


// Read the serialzed blob from the hash
//
// Return the content of the file in string type
func ReadFileBlobWithSerialize(hash string) (string, error) {
	fullObjectsDir := filepath.Join(ObjectsDirName, hash[:2])

	byte, err := os.ReadFile(filepath.Join(fullObjectsDir, hash))
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer(byte)
	blob, err := DeserializeBlob(buff)
	if err != nil {
		return "", err
	}

    return string(blob), nil
}

