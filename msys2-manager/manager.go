package main

import (
   "bufio"
   "errors"
   "github.com/mholt/archiver/v3"
   "io/ioutil"
   "os"
   "path"
   "strings"
)

func baseName(s, char_s string) string {
   n := strings.IndexAny(s, char_s)
   if n == -1 {
      return s
   }
   return s[:n]
}

func getRepo(s string) string {
   if s == "mingw64.db.tar.gz" || strings.HasPrefix(s, "mingw-w64-x86_64-") {
      return "http://repo.msys2.org/mingw/x86_64/"
   }
   return "http://repo.msys2.org/msys/x86_64/"
}

func isFile(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.Mode().IsRegular()
}

func unarchive(in_path, out_path string) error {
   tar_o := &archiver.Tar{OverwriteExisting: true}
   in_file := path.Base(in_path)
   println("EXTRACT", in_file)
   switch path.Ext(in_file) {
   case ".zst":
      zstd_o := archiver.TarZstd{Tar: tar_o}
      return zstd_o.Unarchive(in_path, out_path)
   case ".xz":
      xz_o := archiver.TarXz{Tar: tar_o}
      return xz_o.Unarchive(in_path, out_path)
   default:
      gz_o := archiver.TarGz{Tar: tar_o}
      return gz_o.Unarchive(in_path, out_path)
   }
}

type manager struct {
   cache string
   packages []os.FileInfo
}

func newManager() (manager, error) {
   cache_s, e := os.UserCacheDir()
   if e != nil {
      return Manager{}, e
   }
   msys_s := path.Join(cache_s, "Msys2")
   dir_a, e := ioutil.ReadDir(msys_s)
   if e != nil {
      return Manager{}, e
   }
   db_a := []string{"mingw64.db.tar.gz", "msys.db.tar.gz"}
   for n := range db_a {
      file_s := db_a[n]
      real_s := path.Join(msys_s, file_s)
      if isFile(real_s) {
         continue
      }
      url_s := GetRepo(file_s) + file_s
      e = Copy(url_s, real_s)
      if e != nil {
         return Manager{}, e
      }
      e = Unarchive(real_s, msys_s)
      if e != nil {
         return Manager{}, e
      }
   }
   return Manager{msys_s, dir_a}, nil
}

func (o manager) getName(pack_s string) (string, error) {
   for n := range o.Packages {
      dir_s := o.Packages[n].Name()
      if strings.HasPrefix(dir_s, pack_s + "-") {
         return dir_s, nil
      }
   }
   return "", errors.New(pack_s)
}

func (o Manager) getValue(pack_s, key_s string) ([]string, error) {
   a := []string{}
   name_s, e := o.GetName(pack_s)
   if e != nil {
      return a, e
   }
   real_s := path.Join(o.Cache, name_s, "desc")
   open_o, e := os.Open(real_s)
   if e != nil {
      return a, e
   }
   scan_o := bufio.NewScanner(open_o)
   dep_b := false
   for scan_o.Scan() {
      line_s := scan_o.Text()
      // STATE 2
      if line_s == key_s {
         dep_b = true
         continue
      }
      // STATE 1
      if ! dep_b {
         continue
      }
      // STATE 4
      if line_s == "" {
         break
      }
      // STATE 3
      a = append(a, baseName(line_s, "=>"))
   }
   return a, nil
}

func (o manager) resolve(pack_s string) (map[string]bool, error) {
   pack_m := map[string]bool{}
   for pack_a := []string{pack_s}; len(pack_a) > 0; pack_a = pack_a[1:] {
      pack_s := pack_a[0]
      dep_a, e := o.GetValue(pack_s, "%DEPENDS%")
      if e != nil {
         return pack_m, e
      }
      pack_m[pack_s] = true
      pack_a = append(pack_a, dep_a...)
   }
   return pack_m, nil
}

func (o manager) sync(tar_s string) error {
   open_o, e := os.Open(tar_s)
   if e != nil {
      return e
   }
   scan_o := bufio.NewScanner(open_o)
   for scan_o.Scan() {
      pack_s := scan_o.Text()
      val_a, e := o.GetValue(pack_s, "%FILENAME%")
      if e != nil {
         return e
      }
      file_s := val_a[0]
      real_s := path.Join(o.Cache, file_s)
      if ! isFile(real_s) {
         url_s := GetRepo(file_s) + file_s
         e := Copy(url_s, real_s)
         if e != nil {
            return e
         }
      }
      e = Unarchive(real_s, `C:\msys2`)
      if e != nil {
         return e
      }
   }
   return nil
}
