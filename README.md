## Usage

[![Circle CI](https://circleci.com/gh/higanworks/envmap.svg?style=svg)](https://circleci.com/gh/higanworks/envmap)

```
package main

import (
  "fmt"
  "github.com/higanworks/envmap"
)

func main() {
  envs := envmap.All()
  keys := envmap.ListKeys()

  fmt.Println(envs["HOME"])
  for _, key := range keys {
    fmt.Println(key)
  }

}
```

Outputs..

```
/Users/sawanoboly
rvm_bin_path
TERM_PROGRAM
GEM_HOME
SHELL
TERM
MYVIMRC
IRBRC
...
```

## TODOs

- filter keys for collect by regexp.
