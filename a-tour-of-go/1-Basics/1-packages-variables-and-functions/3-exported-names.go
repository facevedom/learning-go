/* In Go, a name is exported if it begins with a capital letter. For example,
   `Pizza` is an exported name, as is `Pi`, which is exported from the `math`
   package.

   `pizza` and `pi` do not start with a capital letter, so they are not exported
*/
package main

import (
    "fmt"
    "math"
)

func main() {
    /* When exporting a package, you can refer only to its exported names. Any
    "unexported" names are not accessible from outside the package.
    */
    fmt.Println(math.Pi)
}
