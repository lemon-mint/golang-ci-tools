# Golang CI Tools Report

Report generated at: 2022-01-13T00:25:35+09:00

Go version: go1.17.6

golang-ci-tools version: 0.1.0-dev

## staticcheck

```
🎉  No staticcheck errors found!
```

## gosec

```
[gosec] 2022/01/13 00:25:37 Including rules: default
[gosec] 2022/01/13 00:25:38 Excluding rules: default
[gosec] 2022/01/13 00:25:38 Import directory: G:\mega\golang\projectv4\golang-ci-tools
[gosec] 2022/01/13 00:25:39 Checking package: main
[gosec] 2022/01/13 00:25:39 Checking file: G:\mega\golang\projectv4\golang-ci-tools\main.go
[gosec] 2022/01/13 00:25:39 Import directory: G:\mega\golang\projectv4\golang-ci-tools\markdown\table
[gosec] 2022/01/13 00:25:40 Checking package: table
[gosec] 2022/01/13 00:25:40 Checking file: G:\mega\golang\projectv4\golang-ci-tools\markdown\table\table.go
[gosec] 2022/01/13 00:25:40 Import directory: G:\mega\golang\projectv4\golang-ci-tools\modgraph
[gosec] 2022/01/13 00:25:41 Checking package: modgraph
[gosec] 2022/01/13 00:25:41 Checking file: G:\mega\golang\projectv4\golang-ci-tools\modgraph\graph.go
[gosec] 2022/01/13 00:25:41 Import directory: G:\mega\golang\projectv4\golang-ci-tools\staticcheck
[gosec] 2022/01/13 00:25:42 Checking package: staticcheck
[gosec] 2022/01/13 00:25:42 Checking file: G:\mega\golang\projectv4\golang-ci-tools\staticcheck\staticcheck.go
[gosec] 2022/01/13 00:25:42 Import directory: G:\mega\golang\projectv4\golang-ci-tools\gocap
[gosec] 2022/01/13 00:25:44 Checking package: gocap
[gosec] 2022/01/13 00:25:44 Checking file: G:\mega\golang\projectv4\golang-ci-tools\gocap\gocap.go
[gosec] 2022/01/13 00:25:44 Import directory: G:\mega\golang\projectv4\golang-ci-tools\golicenses
[gosec] 2022/01/13 00:25:45 Checking package: golicenses
[gosec] 2022/01/13 00:25:45 Checking file: G:\mega\golang\projectv4\golang-ci-tools\golicenses\licenses.go
[gosec] 2022/01/13 00:25:45 Import directory: G:\mega\golang\projectv4\golang-ci-tools\gosec
[gosec] 2022/01/13 00:25:46 Checking package: gosec
[gosec] 2022/01/13 00:25:46 Checking file: G:\mega\golang\projectv4\golang-ci-tools\gosec\gosec.go
Results:


[1;36mSummary:[0m
  Gosec  : dev
  Files  : 7
  Lines  : 317
  Nosec  : 0
  Issues : [1;32m0[0m


```

## gocap

```
github.com/lemon-mint/golang-ci-tools (file, runtime)

github.com/lemon-mint/golang-ci-tools/gocap (execute, file)
github.com/lemon-mint/golang-ci-tools/golicenses (file, execute)
github.com/lemon-mint/golang-ci-tools/gosec (execute, file)
github.com/lemon-mint/golang-ci-tools/modgraph (execute)
github.com/lemon-mint/golang-ci-tools/staticcheck (execute, file)

```

## go-licenses

| Package Name | License File | License |
| --- | --- | --- |
| [github.com/lemon-mint/golang-ci-tools](https://pkg.go.dev/github.com/lemon-mint/golang-ci-tools) | [https://github.com/lemon-mint/golang-ci-tools/blob/master/LICENSE](https://github.com/lemon-mint/golang-ci-tools/blob/master/LICENSE) | Unlicense |



## Dependencies

Total dependencies: 0
<details><summary>Show Full Dependencies</summary>

</details>

