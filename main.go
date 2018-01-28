package main

import (
  "os"
  "gopkg.in/urfave/cli.v1"
)

func main() {
  cli.NewApp().Run(os.Args)
}
