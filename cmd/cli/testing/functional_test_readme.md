# Functional Test

## How to run

To run all tests, execute following command in this folder:

```sh
go test
# or with verbose -v
go test -v
```

To run a specific test, execute following command in this folder:

```sh
go test github.com/machichima/vcs-go/cmd/cli/testing -v -run {TestFunction}

# e.g.
go test github.com/machichima/vcs-go/cmd/cli/testing -v -run TestRm
```

## How to add functional test cases

Add test case as the txt file with all the commands and their expected outputs.
With `>` represent the input and `<<<` be the separator for the input and output.
The test case file should look like:

```
> input command
<<<
expected output
line 1
line 2
...
> input command 2
<<<
expected output 2
...
```
See the testCases folder for the example test cases.
