package ascii

import "math/rand"


var Cow = `
       \  ^__^
        \ (oo)\_______
          (__)\       )\/\
              ||----w |
              ||     ||
  `

var Stegasaurus = `
         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

`

// GetAnimal: returns a random ascii art animal to speak the phrase
// based on flag passed into script
// if nothing is asked for, sends random animal
func GetAnimal(flag string) string  {

  var rnd = rand.Intn(2)

  switch flag {
  case "cow":
    return Cow

  case "stego":
    return Stegasaurus

  default:
    switch rnd {
    case 1:
      return Cow
    case 2:
      return Stegasaurus
    default:
      return  Cow
      
    }
  }
}
