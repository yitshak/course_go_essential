package main

import (
  "io"
  "os"
  "fmt"
)

type Capper struct{
  wtr io.Writer
}

func (capper Capper) Write(p []byte)(n int , err error){
  for i:=0; i<len(p);i++{
    if p[i]>=byte('a') && p[i] <=byte('z'){
      p[i] -= 32
    }
  }
  return capper.wtr.Write(p)
}

func main() {
  testString:= "thiS is a SUPEr duper 555 (*)&*"
  capper := &Capper{os.Stdout}
  fmt.Fprintln(capper, testString)
}