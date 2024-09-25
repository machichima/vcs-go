package utils

// Combine the index file and previous filetree
func MergeIndexAndFileTree(index Index, fileTree Index) Index {
    // loop through index file and update the previous filetree
    for file, hash := range index.FileToHash {
        if hash == "" {
            // delete the file
            delete(fileTree.FileToHash, file)
            continue
        }

        // update the file hash
        fileTree.FileToHash[file] = hash
    }

    return fileTree
}
