# Image Server
An image server toolkit in Go (Golang)

## Features
- Http server
- Resize / convert / process (Graphicsmagick)
- Cache (Redis, Memcache, in memory)
- Fully modular

## Demo
```
Normal
http://fuckingfrogs.fr:8080/?source=small.jpg
```
![Normal](http://fuckingfrogs.fr:8080/?source=small.jpg)

```
Resize animated gif
http://fuckingfrogs.fr:8080/?source=animated.gif&width=300&height=300
```
![Resize animated gif](http://fuckingfrogs.fr:8080/?source=animated.gif&width=300&height=300)

```
Resize and crop
http://fuckingfrogs.fr:8080/?source=medium.jpg&width=200&height=200&extent=1&fill=1
```
![Resize and crop](http://fuckingfrogs.fr:8080/?source=medium.jpg&width=200&height=200&extent=1&fill=1)

```
Resize jpeg low quality
http://fuckingfrogs.fr:8080/?source=large.jpg&width=400&format=jpeg&quality=50
```
![Resize jpeg low quality](http://fuckingfrogs.fr:8080/?source=large.jpg&width=400&format=jpeg&quality=50)

```
Resize huge image (5000x5000)
http://fuckingfrogs.fr:8080/?source=huge.jpg&width=300&height=300
```
![Resize huge image (5000x5000)](http://fuckingfrogs.fr:8080/?source=huge.jpg&width=300&height=300)

## Build status
[![Build Status](https://travis-ci.org/pierrre/imageserver.png?branch=master)](https://travis-ci.org/pierrre/imageserver)

## Examples
- [Simple](https://github.com/pierrre/imageserver/blob/master/_examples/simple/simple.go)
- [Advanced](https://github.com/pierrre/imageserver/blob/master/_examples/advanced/advanced.go)

## Documentation
- http://godoc.org/github.com/pierrre/imageserver
- http://sourcegraph.com/github.com/pierrre/imageserver

## Help
- Twitter: @pierredurand87
- Github issue

## TODO
- more tests
- don't ignore error from cache
- add GroupcacheServer
- add timeout in LimitProcessor
