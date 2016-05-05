# nettop
a simple command line Bandwidth Monitor by parsing `/proc/net/dev` to calculator.

## Build
```
$ go build -o nettop nettop.go
```

## Usage
```
Usage of nettop:
  -c uint
    	count (0 == unlimit)
  -i string
    	interface (default "*")
  -t float
    	update time(s) (default 2)
  -v int
    	verbosity (default 2)

```

## Output
```
$ ./nettop -i eth0 -t 1 -c 3
BW:	eth0	66.00 B/s	66.00 B/s
BW:	eth0	157.00 B/s	143.00 B/s
BW:	eth0	163.00 B/s	224.00 B/s
```

## LICENSE - MIT
Copyright (C) 2016 cs8425

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


