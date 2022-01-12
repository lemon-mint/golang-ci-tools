# Golang CI Tools Report

Report generated at: 2022-01-13T07:11:58+09:00

Go version: go1.17.6

golang-ci-tools version: 0.1.1-dev

## staticcheck

```
🎉  No staticcheck errors found!
```

## gosec

```
[gosec] 2022/01/13 07:12:00 Including rules: default
[gosec] 2022/01/13 07:12:00 Excluding rules: default
[gosec] 2022/01/13 07:12:00 Import directory: G:\mega\golang\projectv4\golang-ci-tools
[gosec] 2022/01/13 07:12:02 Checking package: main
[gosec] 2022/01/13 07:12:02 Checking file: G:\mega\golang\projectv4\golang-ci-tools\main.go
[gosec] 2022/01/13 07:12:02 Import directory: G:\mega\golang\projectv4\golang-ci-tools\markdown\table
[gosec] 2022/01/13 07:12:03 Checking package: table
[gosec] 2022/01/13 07:12:03 Checking file: G:\mega\golang\projectv4\golang-ci-tools\markdown\table\table.go
[gosec] 2022/01/13 07:12:03 Import directory: G:\mega\golang\projectv4\golang-ci-tools\modgraph
[gosec] 2022/01/13 07:12:04 Checking package: modgraph
[gosec] 2022/01/13 07:12:04 Checking file: G:\mega\golang\projectv4\golang-ci-tools\modgraph\graph.go
[gosec] 2022/01/13 07:12:04 Import directory: G:\mega\golang\projectv4\golang-ci-tools\staticcheck
[gosec] 2022/01/13 07:12:05 Checking package: staticcheck
[gosec] 2022/01/13 07:12:05 Checking file: G:\mega\golang\projectv4\golang-ci-tools\staticcheck\staticcheck.go
[gosec] 2022/01/13 07:12:05 Import directory: G:\mega\golang\projectv4\golang-ci-tools\gocap
[gosec] 2022/01/13 07:12:07 Checking package: gocap
[gosec] 2022/01/13 07:12:07 Checking file: G:\mega\golang\projectv4\golang-ci-tools\gocap\gocap.go
[gosec] 2022/01/13 07:12:07 Import directory: G:\mega\golang\projectv4\golang-ci-tools\golicenses
[gosec] 2022/01/13 07:12:08 Checking package: golicenses
[gosec] 2022/01/13 07:12:08 Checking file: G:\mega\golang\projectv4\golang-ci-tools\golicenses\licenses.go
[gosec] 2022/01/13 07:12:08 Import directory: G:\mega\golang\projectv4\golang-ci-tools\gosec
[gosec] 2022/01/13 07:12:09 Checking package: gosec
[gosec] 2022/01/13 07:12:09 Checking file: G:\mega\golang\projectv4\golang-ci-tools\gosec\gosec.go
Results:


Summary:
  Gosec  : dev
  Files  : 7
  Lines  : 319
  Nosec  : 0
  Issues : 0


```

## gocap

```
github.com/lemon-mint/golang-ci-tools (file)

github.com/lemon-mint/golang-ci-tools/gocap (execute, file)
github.com/lemon-mint/golang-ci-tools/golicenses (execute, file)
github.com/lemon-mint/golang-ci-tools/gosec (execute, file)
github.com/lemon-mint/golang-ci-tools/modgraph (execute)
github.com/lemon-mint/golang-ci-tools/staticcheck (execute, file)

```

## go-licenses

| Package Name | License File | License |
| --- | --- | --- |
| [github.com/lemon-mint/golang-ci-tools](https://pkg.go.dev/github.com/lemon-mint/golang-ci-tools) | [https://github.com/lemon-mint/golang-ci-tools/blob/master/LICENSE](https://github.com/lemon-mint/golang-ci-tools/blob/master/LICENSE) | Unlicense |
| [github.com/acarl005/stripansi](https://pkg.go.dev/github.com/acarl005/stripansi) | [https://github.com/acarl005/stripansi/blob/master/LICENSE](https://github.com/acarl005/stripansi/blob/master/LICENSE) | MIT |



## Dependencies

Total dependencies: 1
<details><summary>Show Full Dependencies</summary>

 - github.com/lemon-mint/golang-ci-tools github.com/acarl005/stripansi@v0.0.0-20180116102854-5a71ef0e047d
</details>

