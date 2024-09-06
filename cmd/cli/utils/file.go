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
func SaveFileByHash(filePath string, hash string, blob []byte, commandType int) error {

	// create parent dir
	parentDir := hash[:2]
	fullObjectsDir := ObjectsDirName + "/" + parentDir

	if err := CreateOneDir(fullObjectsDir); err != nil {
		return err
	}

	// write blob to file
	if err := os.WriteFile(fullObjectsDir+"/"+hash, blob, os.ModePerm); err != nil {
		return err
	}

	if commandType == AddType {
		//read index
		index, err := ReadIndexFile()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				// if the index file does not exist, create one with empty index struct
				fmt.Println("Index file does not exist")
                index = Index{FileToHash: make(map[string]string)}
			} else {
				return err
			}
		}
		if err := AddToIndex(&index, filePath, hash); err != nil {
			return err
		}

        if err := WriteIndexFile(index); err != nil {
            return err
        }
	}

	return nil
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
func AddToIndex(index *Index, file string, hash string) error {
	// file already exists in index
	if index.FileToHash[file] != "" && index.FileToHash[file] != hash {
		if err := DeleteObject(index.FileToHash[file]); err != nil {
			return err
		}
	}

	// update or add file-hash to index
	index.FileToHash[file] = hash

	return nil
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
