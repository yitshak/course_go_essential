// Get content type of sites
package main

import (
	"fmt"
  "os"
  "github.com/pkg/errors"
  "bufio"
  "strings"
  "crypto/md5"
  "io"
)

func parseSignatureFile( filePath string) (map[string]string,error){
  file, err := os.Open(filePath)
  if err != nil{
    return nil, errors.Wrap(err,"ERROR: problem opening signatures file")
  }
  defer file.Close()

  sigs := make(map[string]string)
  scanner := bufio.NewScanner(file)

  for lnum := 1;scanner.Scan(); lnum++{
    //6c6427da7893932731901035edbb9214  nasa-00.log
    fields := strings.Fields(scanner.Text())
    if len(fields) !=2 {
      return nil,fmt.Errorf("ERROR: %v: Problem parsing line %v \n",filePath, lnum)
    }
    sigs[fields[1]] = fields[0]
  }

  if err := scanner.Err(); err != nil {
    return nil,errors.Wrap(err,"ERROR: reading signature file\n")
  }

  return sigs, nil
}

func fileMD5(filePath string)(string, error){
  file, err := os.Open("nasa-log/"+filePath)
  if err != nil{
    return "",errors.Wrap(err,"ERROR: problem opening file for MD5 calculation")
  }
  defer file.Close()
  hash := md5.New()

  if _, err = io.Copy(hash, file); err != nil{
    return "", errors.Wrap(err,"ERROR: calculating md5")
  }

  return fmt.Sprintf("%x",hash.Sum((nil))), nil
}

type result struct{
  path string
  match bool
  err error
}

func md5Worker(path string, signature string, out chan *result){
  r:= &result{path: path}
  calculatedSig, err := fileMD5(path)
  r.err = err
  if err != nil {
    out <- r
    return
  }
  r.match = signature == calculatedSig
  out <- r
  return

}

func main() {

  fileMap, err := parseSignatureFile("nasa-log/md5sum.txt")
	if err != nil{
    fmt.Fprintf(os.Stderr,"%v", err)
  }

  ch := make(chan *result)

	for filePath, signature := range fileMap {
		go md5Worker(filePath, signature, ch)
	}

  for range fileMap{
    result := <-ch
    if result.err != nil{
      fmt.Fprintf(os.Stderr,"%v\n", result.err)
    }else{
      if result.match{
        fmt.Printf("MD5 signature for %v - match\n", result.path)
      }else{
        fmt.Printf("MD5 signature for %v - do not match\n", result.path)
      }
      
    }
		
	}
  close(ch)
}
