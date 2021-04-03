package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "log"
   "os"
   "os/exec"
   "strings"
   "time"
)

const minimum = 64
var add, del, totAdd, totCha, totDel int

func diff() (*bufio.Scanner, error) {
   _, e := os.Stat("config.toml")
   if e != nil {
      return popen("git", "diff", "--cached", "--numstat")
   }
   return popen("git", "diff", "--cached", "--numstat", ":!docs")
}

func popen(name string, arg ...string) (*bufio.Scanner, error) {
   cmd := exec.Command(name, arg...)
   pipe, e := cmd.StdoutPipe()
   if e != nil {
      return nil, fmt.Errorf("StdoutPipe %v", e)
   }
   return bufio.NewScanner(pipe), cmd.Start()
}

type test struct {
   name string
   actual, target interface{}
   result bool
}

func main() {
   c := exec.Command("git", "add", ".")
   c.Stderr, c.Stdout = os.Stderr, os.Stdout
   e := c.Run()
   if e != nil {
      log.Fatal(e)
   }
   stat, e := diff()
   if e != nil {
      log.Fatal(e)
   }
   for stat.Scan() {
      totCha++
      text := stat.Text()
      if strings.HasPrefix(text, "-") {
         continue
      }
      fmt.Sscanf(text, "%v\t%v", &add, &del)
      totAdd += add
      totDel += del
   }
   commit, e := popen("git", "log", "--format=%cI")
   if e != nil {
      log.Fatal(e)
   }
   commit.Scan()
   // actual
   actual := commit.Text()[:10]
   // target
   target := time.Now().AddDate(0, 0, -1).String()[:10]
   // print
   for _, each := range []test{
      {"additions", totAdd, minimum, totAdd >= minimum},
      {"deletions", totDel, minimum, totDel >= minimum},
      {"changed files", totCha, minimum, totCha >= minimum},
      {"last commit", actual, target, actual <= target},
   } {
      message := fmt.Sprintf(
         "%-16v target: %-12v actual: %v", each.name, each.target, each.actual,
      )
      if each.result {
         x.LogPass("Pass", message)
      } else {
         x.LogFail("Fail", message)
      }
   }
}
