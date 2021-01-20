package main

import (
   "github.com/89z/json"
   "golang.org/x/build/repos"
   "sort"
)

func count() error {
   var repo_a []repo
   for repo_s, repo_o := range repos.ByImportPath {
      if ! repo_o.ShowOnDashboard() {
         continue
      }
      m, e := json.LoadHttp("https://api.godoc.org/search?q=" + repo_s + "/")
      if e != nil {
         return e
      }
      results := m.A("results")
      size := len(results)
      repo_a = append(repo_a, repo{size, repo_s})
   }
   sort.Slice(repo_a, func(n, n2 int) bool {
      return repo_a[n].count < repo_a[n2].count
   })
   for _, o := range repo_a {
      println(o.count, o.path)
   }
   return nil
}

type repo struct {
   count int
   path string
}