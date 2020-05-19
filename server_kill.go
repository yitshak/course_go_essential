package main

import (
  "fmt"
  "github.com/pkg/errors"
  "io/ioutil"
  "strconv"
  "os"
  "log"
)

func killServer(pidFile string) error {
  
  defer func() {
    if err := recover(); err != nil {
      fmt.Fprintf(os.Stderr, "ERROR in killServer: %v",err)
    }
  }()

  pidString, err := ioutil.ReadFile(pidFile)
  if err != nil {
    return errors.Wrap(err,"Cannot open pidFile: ")
  }
  pid,err := strconv.Atoi(string(pidString))
  if err != nil {
    return errors.Wrap(err,"Cannot convert string to pid number ")
  }
  fmt.Printf("Killing pid %v\n", pid)
  return nil
}

func setupLoging() {
  logFile, err := os.OpenFile("ServerKill.log",
                              os.O_APPEND | os.O_CREATE |os.O_WRONLY, 0644)
  if err != nil{
    fmt.Fprintf(os.Stderr,"%v\n",errors.Wrap(err,"Problem with logfile:"))
  }
  log.SetOutput(logFile)
}

func main() {
  setupLoging()
  err := killServer("no-such-file")
  fmt.Fprintf(os.Stderr, "No such file case:\n %v\n", err)
  log.Printf("%+v\n",err)

  err = killServer("no_pid.txt")
  fmt.Fprintf(os.Stderr,"No PID in text:\n %v\n", err)
  log.Printf("%+v\n",err)

	err = killServer("pid.txt")
  fmt.Fprintf(os.Stderr,"PID in text:\n %v\n", err)
  log.Printf("%+v\n",err)
}
