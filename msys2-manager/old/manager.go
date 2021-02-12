package main

import (
   "bufio"
   "fmt"
   "github.com/89z/x"
   "github.com/89z/x/extract"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

type manager struct {
   x.Install
}

func baseName(s, char string) string {
   n := strings.IndexAny(s, char)
   if n == -1 {
      return s
   }
   return s[:n]
}

func (m manager) sync(tar string) error {
   open, e := os.Open(tar)
   if e != nil {
      return e
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      values, e := m.getValue(
         scan.Text(), "%FILENAME%",
      )
      if e != nil {
         return e
      }
      file := values[0]
      archive := path.Join(m.Cache, file)
      _, e = x.Copy(
         getRepo(file) + file, archive,
      )
      if e != nil {
         return e
      }
      e = unarchive(archive, m.Dest)
      if e != nil {
         return e
      }
   }
   return nil
}

func main() {
   man := manager{install}
   target := os.Args[2]
   if os.Args[1] == "sync" {
      e = man.sync(target)
      x.Check(e)
      return
   }
   packSet := map[string]bool{}
   for packs := []string{target}; len(packs) > 0; packs = packs[1:] {
      target := packs[0]
      deps, e := man.getValue(target, "%DEPENDS%")
      x.Check(e)
      packs = append(packs, deps...)
      if packSet[target] {
         continue
      }
      println(target)
      packSet[target] = true
   }
}
