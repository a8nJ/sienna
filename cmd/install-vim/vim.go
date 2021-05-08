package main

import (
   "github.com/89z/rosso"
   "os"
   "path"
)

var runtime = []struct{dir, base string}{
   {
      "sienna/vim/colors",
      "NLKNguyen/papercolor-theme/e397d18a/colors/PaperColor.vim",
   }, {
      "sienna/vim/ftdetect",
      "zah/nim.vim/master/ftdetect/nim.vim",
   }, {
      "sienna/vim/ftdetect",
      "PProvost/vim-ps1/master/ftdetect/ps1.vim",
   }, {
      "sienna/vim/syntax",
      "dart-lang/dart-vim-plugin/master/syntax/dart.vim",
   }, {
      "sienna/vim/syntax",
      "google/vim-ft-go/master/syntax/go.vim",
   }, {
      "sienna/vim/syntax",
      "vim/vim/a942f9ad/runtime/syntax/javascript.vim",
   }, {
      "sienna/vim/syntax",
      "tpope/vim-markdown/564d7436/syntax/markdown.vim",
   }, {
      "sienna/vim/syntax",
      "zah/nim.vim/master/syntax/nim.vim",
   }, {
      "sienna/vim/syntax",
      "PProvost/vim-ps1/master/syntax/ps1.vim",
   }, {
      "sienna/vim/syntax",
      "vim/vim/b9c8312e/runtime/syntax/python.vim",
   }, {
      "sienna/vim/syntax",
      "cespare/vim-toml/master/syntax/toml.vim",
   },
}

func main() {
   zip := path.Join(
      "vim",
      "vim-win32-installer",
      "releases",
      "download",
      "v8.2.2677",
      "gvim_8.2.2677_x64.zip",
   )
   inst := rosso.NewInstall("sienna/vim", zip)
   inst.SetCache()
   _, err := rosso.Copy("https://github.com/" + zip, inst.Cache)
   if os.IsExist(err) {
      println("Exist", inst.Cache)
   } else if err != nil {
      panic(err)
   }
   arc := rosso.Archive{2}
   println("Zip", inst.Cache)
   arc.Zip(inst.Cache, inst.Dest)
   for _, each := range runtime {
      inst = rosso.NewInstall(each.dir, each.base)
      os.Remove(inst.Dest)
      _, err = rosso.Copy("https://raw.githubusercontent.com/" + each.base, inst.Dest)
      if err != nil {
         panic(err)
      }
   }
}
