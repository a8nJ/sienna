package main
import "os"

func main() {
   if len(os.Args) != 2 {
      println("rust-deps <crate>")
      os.Exit(1)
   }
   crate := os.Args[1]
   e := system("cargo", "new", "rust-deps")
   check(e)
   os.Chdir("rust-deps")
   toml := Map{
      "dependencies": Map{crate: ""},
      "package": Map{"name": "rust-deps", "version": "1.0.0"},
   }
   e = tomlEncode("Cargo.toml", toml)
   check(e)
   e = system("cargo", "generate-lockfile")
   check(e)
   lock, e := tomlDecode("Cargo.lock")
   check(e)
   prev := ""
   dep := 0
   packages := lock.A("package")
   for n := range packages {
      name := packages.M(n).S("name")
      if name == "rust-deps" {
         continue
      }
      if name == crate {
         continue
      }
      if name == prev {
         continue
      }
      println(name)
      prev = name
      dep++
   }
   print("\n", dep, " deps\n")
   os.Chdir("..")
   e = os.RemoveAll("rust-deps")
   check(e)
}
