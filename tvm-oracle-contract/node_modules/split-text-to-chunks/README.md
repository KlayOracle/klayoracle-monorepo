# Split text into chunks

Prefers splitting at whitespace characters, but falls back to "hard wrapping",
and obeys existing newlines.

```js
const {split,width} = require('split-text-to-chunks')

const str = 'A lazy 🐶 made a pile of 💩'

split(str, 8/*columns*/)
// -> [ 'A lazy 🐶', 'made a ', 'pile of ', '💩' ]

split('A lazy dog', 3)
// -> [ 'A ', 'laz', 'y ', 'dog' ]

split('A lazy\ndog', 10)
// -> [ 'A lazy', 'dog' ]

width(str)
// -> 25

width(str, 20/* max, stop counting */)
// -> 20

width('one\ntwo\nthree')
// -> 5
```

```sh
$ npm i -g split-text-to-chunks

$ printf "A lazy 🐶 made a pile of 💩" | wordwrap --columns 8 # default: 80
A lazy 🐶
made a
pile of
💩

$ printf "A lazy 🐶\nmade a pile of 💩" | wordwrap --width
16
```
