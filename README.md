# jscrush

This is a command line version of [jscrush][1]. I made it to compress my entry
for the [js1k.com][2] contest. It does not strip comments nor white spaces, it
expects an output from [uglifyjs][3]. Does not have options or any cool
features, it's just a quick and dirty tool.

## Install

```
$ go get github.com/xiam/jscrush
```

## Usage

```
$ cat script.js | uglifyjs | jscrush > crushed.js
```

[1]: http://www.iteral.com/jscrush/
[2]: http://www.js1k.com
[3]: https://github.com/mishoo/UglifyJS
