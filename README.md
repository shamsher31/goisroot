# goisroot

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goisroot)
[![Build Status](https://travis-ci.org/shamsher31/goisroot.svg)](https://travis-ci.org/shamsher31/goisroot)
[![GitHub release](http://img.shields.io/github/release/shamsher31/goisroot.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)

Check if the process is running as root user, eg. started with sudo

### How to install
```go
go get github.com/shamsher31/goisroot
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goisroot"
)

func main() {
  fmt.Println(root.Is())
  // true
}

```
###Aliasing Imports
If you already have package name ```root``` you can use following.
```go
package main

import (
	"fmt"
	pckRoot "github.com/shamsher31/goisroot"
)

func main() {
  fmt.Println(pckRoot.Is())
  // true
}
```

### Related
[goisdocker](https://github.com/shamsher31/goisdocker)<br>
[goistravis](https://github.com/shamsher31/goistravis)<br>
[goisredirect](https://github.com/shamsher31/goisredirect)<br>

### Why
This package is inspired by [is-root](https://www.npmjs.com/package/is-root) npm module.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
