# Go Cowsay

a clone of the hilarious linux cli classic `cowsay`

## Install via Go

`go get github.com/egotch/gocowsay`

```bash
go install github.com/egotch/gocowsay
```

## Usage

* only accepts piped input (i.e. `fortune | gocowsay`)
* can specify the figure to read back text
    * -f arg specifies the figure (cow, stego, rnd)

## Examples


```bash
fortune | gocowsay -f stego
```

```
 -----------------------------------------
/                                         \
| Exercise caution in your daily affairs. |
\                                         /
 -----------------------------------------

         \                      .       .
          \                    / `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    `._.'               |         }  \/~~~/
       `----.          /       }     |        /    \__/
             `-.      |       /      |       /      `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      `..-'
                      |     /        |     /     ~-.     `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

```
