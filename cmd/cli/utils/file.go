package utils

import (
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
// of the hash
//
// e.g. hash1: dfq8hfroihffjlkasj / hash2: df32fjqf81efh1ofj
// saved in df/dfq8hfroihffjlkasj and df/df32fjqf81efh1ofj file
func SaveFileByHash(hash string, blob []byte, commandType int) error {
	// TODO: implement this function

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
		// TODO: write file name and hash to index file
	}

	return nil
}


// TODO: add test (after the test for DeleteObject
// Add to INDEX file
func AddToIndex(index *Index, file string, hash string) error {
    // file already exists in index
    if index.FileToHash[file] != "" && index.FileToHash[file] != hash {
        // TODO: delete original objects files
        if err := DeleteObject(index.FileToHash[file]); err != nil {
            return err
        }
    }

    // update or add file-hash to index
    index.FileToHash[file] = hash

    return nil
}


// TODO: add test for this
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
