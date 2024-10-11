# vcs-go


## Introduction

> This is the Go version of Gitlet from [this project](https://sp21.datastructur.es/materials/proj/proj2/proj2) written in JAVA. I try to follow all the concept in this project and reproduce in Go.

- store files / metadata in `.vgo/` folder

- commands
    - `init`
    - `add`
    - `status`
    - `commit`
    - `rm`
        - unstage the file
    - `log`
        - show the history of current node
    - `checkout`
        - file: take out the file from the HEAD commit to the working directory
        - commit_id file: take out the file from the specific commit to the working directory
        - (advanced) branch_name: switch to the branch
    - (advanced) other commands for branch -> refer to external link for detail


## file structure

- File structure for the project is as follows:
```{sh}
.
├── cmd
│   └── cli
│       ├── commands
│       │   ├── init.go
│       │   └── root.go
│       ├── main.go
│       └── utils
│           ├── file.go
│           └── serialize.go
├── go.mod
├── go.sum
├── README.md
└── testfolder
```

Test for each file are included in the same folder as the file itself.

- `utils`: tool functions
- `commands`: each commands shown in [Introduction](##Introduction) has its own folder


- The committed files are stored in `.vgo/objects` folder. The file structure is as follows:
```{sh}
/my_project
    .vgo
        /objects
            6a/
                7f3  # Serialized content of a blob with hash starting with 6a7f3...
            2f/
                4f2  # Serialized content of a tree with hash starting with 2f4f2...
            d1/
                2f5  # Serialized content of a commit with hash starting with d12f5...
        /refs
            /heads
                main  # Pointer to the latest commit on the main branch
        HEAD  # Pointer to the current commit checked out
```

## Usage

### Run

- To run the project, run following in the target folder (testfolder for testing our project):

```{sh}
go run ../cmd/cli/main.go {command}

```

### Test

- Go to the folder of the file you want to test, run following:
```{sh}
go test
```
To see more details of the test, run:
```{sh}
go test -v
```

To test specific function in a module, run:
```{sh}
go test github.com/machichima/vcs-go/cmd/cli/utils -v -run function
```


## TODO

- [x] init
    - [x] create .vcsgo folder with empty structures inside
    - [x] Detect if .vcsgo folder already exists
- [x] Serialization
    - [x] files into blob
    - [x] Serialize the file blob and deserialize
    - [x] Hash the file blob
    - [x] file tree into blob
        - [x] Get file with directory structure
        - [x] hashmap for file dir structure to SHA-1 hash of the file
    - [x] commit into blob
- [x] SHA-1 hashing function
    - [x] Hash from serialized objects -> get hash string
    - [x] Store hash string in .vcsgo/objects
- [x] Track difference between current files and previous commit
    - used in stage function
- [x] Combine the filetree from new and old commit to form a new one
    - always save the newest full filetree of the project

- [x] add
- [x] status
- [x] commit
- [x] log
- [x] Write test (functional test)
    - [x] add
    - [x] status
    - [x] commit
    - [x] log
- [x] rm
- [x] checkout
- [ ] branch
- [ ] branch -d (delete branch)
- [ ] merge
- [ ] remote option ...

## Work Log

### 11/10/2024
- [x] Finish tests for `checkout`

### 07/10/2024
- [x] finish functional test for add / status / rm / commit / log commands

### 26/09/2024
- [x] finish rm command

### 25/09/2-24
- [x] Finish until log function
    - update filetree based on index and prev filetree

### 13/09/2024
- [x] Track difference between current files and previous commit
    - used in stage function
- [x] Combine the filetree from new and old commit to form a new one
    - always save the newest full filetree of the project

### 12/09/2024
- [x] Finish status command
- [x] Finish commit command
- [x] Finish log command

### Before 12/09/2024
- [x] Finish the add function

### 27/07/2024
- [x] Finish Serialization function for the file contents
- [x] Get the files with their directory structure
