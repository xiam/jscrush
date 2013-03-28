# jscrush

This is a command line version of [jscrush][1]. I made it just to compress my
entry ("[Have you ever tried to catch a firefly?][4]") for the [js1k.com][2]
contest. It does not strip comments nor white spaces, it just searches for
string repetitions, so in order to be useful you can feed it with the output
of [uglifyjs][3]. Does not have options or any cool features, it's just a quick
and dirty tool that is not even optimized or beauty.

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
[4]: http://js1k.com/2013-spring/demo/1462
