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

## TODO

- [x] init
    - [x] create .vcsgo folder with empty structures inside
    - [x] Detect if .vcsgo folder already exists
- [ ] Serialization
    - [x] files into blob
        - [ ] Hash the file blob
    - [ ] file tree into blob
        - [x] Get file with directory structure
        - [ ] hashmap for file dir structure to SHA-1 hash of the file
    - [ ] commit into blob
- [ ] SHA-1 hashing function
    - [ ] Hash from serialized objects -> get hash string
    - [ ] Store hash string in .vcsgo/objects
- [ ] add
- [ ] commit
- [ ] log
- [ ] checkout
- [ ] rm
- [ ] branch operation...


## Work Log

### 27/07/2024
- [x] Finish Serialization function for the file contents
- [x] Get the files with their directory structure
