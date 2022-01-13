# ds

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

List es una librería para gestionar una Lista enlazada.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales

> Objeto List
- *tipo* `LinkedList`, con métodos:  
- `Add(...Item)`
- `Index(Item) int`
- `Slice() []interface{}`
- `IsEmpty() bool`
- `Len() int` 
- `String() string `
- `Insert(int, Item) error`
- `Remove()` 
- `RemoveAt(int)(*Item,error)` 
- `Contains(Item) bool`

> Funciones
- `NewList(...Item)`

> Eventos 
- *`Init(*List) error {}`* 					in 	NewList
- *`Insert(*List, *Node) error {}`* 	in Insert
-	*`Get(*List, *Node) error {}`* in	Get
-	*`Remove(*List, *Node) error {}`* in Remove
-	*`Pop(*List, *Node) error {}`* in 	Pop
- *`Move(*List, int) error {}`* in	MoveTo


## Ejemplos
```go

	list := NewList()
	flaga := list.Empty()
	list.Push("solo", "han", 16, "norris")
	node := list.Head.next // obtiene iterador
	flaga = list.Empty()
	fore := list.Slice()
	list.Push("jesus")
	list.Push(25)
	list.Push("torres")
	list.Push("sacristan")
	foreA := list.Slice()[2]
	flag := list.Contains("jesus")
	nb := list.Length()
	list.Insert(0, "gilipollas")
	list.Insert(1, "probaperas")
	list.Insert(4, "tocomocho")
	fore = list.Slice()
	negre, _ := list.Pop(2)
	fore = list.Slice()
	negre, _ = list.Pop(0)
	fore = list.Slice()
	negre, _ = list.Pop(6)
	fore = list.Slice()
	list.Push("jesus")
	list.Push("torres")
	list.Push("sacristan")
	fore = list.Slice()
	nb = list.Length()
	list.Remove()

	fmt.Println(flag, flaga, fore, foreA, negre, nb, node)

```
## Notas




<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
