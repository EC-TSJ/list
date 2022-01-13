/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/
package list

import (
	"ec-tsj/core"
	"ec-tsj/event"
	"fmt"
)

// T the type of the linked list
type (
	T = core.T

	Paragraph struct {
		Value T
		Info  string
	}

	// Node a single Node that composes the list
	Node struct {
		info  string
		value T
		next  *Node
	}

	// List the linked list of Items
	List struct {
		Head *Node // nodo centinela
		size int
	}

	fSet func() *List

	// Equivalent for Paragraph
	P = Paragraph
)

var (
	Events event.IEventEmitter
)

/*
* inicializa datos de settings
 */
func init() {
	Events = event.EventEmitter("<List>", false)
}

/**
 * Devuelve el valor del Node
 * return {T}
 */
func (node *Node) Value() T {
	return node.value
}

/**
 * Avanza al próximo Node
 * @return {*Node}
 */
func (node *Node) Next() *Node {
	return node.next
}

/**
 * Info extra
 * @return {string}
 */
func (node *Node) Info() string {
	return node.info
}

var _S_OUT_ string = "** Índice fuera de límites **"

// helper para remove y NewList
func __set() *List {
	return &List{Head: &Node{value: "@Init@", next: nil, info: "Elemento nulo."}, size: 0}
}

// Equivalent of __set()
var Sx fSet = __set

/**
 * NewList crea una nueva List, con varios items o ninguno.
 * NewList().Head es un iterador.
 * NewList().Head.next es el iterador (tb) correspondiente al primer
 * item de la sucesión.
 * <CODE>
 *	...
 * 	list := NewList()  				// list is List
 * 	node := list.Head.next     // node is a Node iterator
 *	...
 * </CODE>
 * <Events>
 * 		"Init", 	func(*List) error {} 					...in... 	NewList
 * 		"Insert", func(*List, *Node) error {} 	... 			Insert
 *		"Get", 		func(*List, *Node) error {} 	... 			Get
 *		"Remove", func(*List, *Node) error {} 	... 			Remove
 *		"Pop",  	func(*List, *Node) error {} 	...				Pop
 *    "Move",  	func(*List, int) error {} 	 	... 			MoveTo
 * </Events>
 * @return {*List}
 */
func NewList(items ...T) *List {
	list := __set()
	Events.Emit("Init", list)
	s := len(items)
	if s != 0 {
		list.Push(items...)
	}
	return list
}

/**
 * Recibe un array con los valores del Stack/Queue
 * @return {[]interface{}}
 */
func (list *List) Slice() []T {
	items := make([]T, 0)
	// apunta después que '@Init@'
	for node := list.Head.next; node != nil; node = node.Next() {
		items = append(items, Paragraph{node.Value(), node.Info()})
	}
	return items
}

/**
 * Push añade un T, o varios, al final de la lista. Base 0
 * @param {...T}
 */
func (list *List) Push(t ...T) {
	for _, f := range t {
		list.Insert(list.size, f)
	}
}

/**
 * Insert añade un T t en la posición i. Base 0
 * @param {int}
 * @param {T}
 * @return {error}
 */
func (list *List) Insert(
	i int,
	t T) error {
	err := list._withinRange(i)
	if err == nil {
		var addnode *Node
		switch g := t.(type) {
		case Paragraph:
			addnode = &Node{value: g.Value, next: nil, info: g.Info}
		default:
			addnode = &Node{value: t, next: nil, info: core.Literals().NullString}
		}
		node := list.Locate(i)
		addnode.next = node.Next()
		node.next = addnode
		list.size++
		Events.Emit("Insert", list, addnode)
	}
	return err
}

/**
 * Localiza un sitio determinado. El indicado por i. Base 0
 * @param {int}
 * @return {*Node}
 */
func (list *List) Locate(i int) *Node {
	node := list.Head
	for j := 0; j < i; node, j = node.Next(), j+1 {
	}
	return node
}

/**
 * Index retorna la posición del T t. Base 0
 * @param {T}
 * @return {int}
 */
func (list *List) Index(t T) int {
	for node, j := list.Head, -1; ; node, j = node.Next(), j+1 {
		if node.Value() == t {
			return j
		}
		if node.Next() == nil {
			return -1
		}
	}
}

/**
 * Obtiene un T de la lista. Base 0
 * @param {int}
 * @return {*Node}
 */
func (list *List) Get(i int) T {
	node := list.Locate(i)
	Events.Emit("Get", list, node)
	return core.IIf(node.next != nil, node.Next().Value(), core.Literals().NullString)
}

/**
 * Pone un T en la lista. Base 0
 * @param {int}
 * @return {*Node}
 */
func (list *List) Set(
	i int,
	t T) {
	list.Insert(i, t)
}

/**
 * IsEmpty retorna true si la lista esta vacía
 * @return {bool}
 */
func (list *List) Empty() bool {
	if list.Head.next == nil {
		return true
	}
	return false
}

/**
 * Len retorna el tamaño de la lista enlazada
 * @return {int}
 */
func (list *List) Length() int {
	return list.size
}

// Interface Stringer
func (list *List) String() string {
	var values []T
	// apunta al posterior a "@Init@"
	for node := list.Head.next; node != nil; node = node.Next() {
		values = append(values, Paragraph{node.Value(), node.Info()})
	}
	return fmt.Sprint(values)
}

// helper function para Pop e Insert
func (list *List) _withinRange(v int) error {
	if v < 0 || v > list.size {
		return core.CustomError("ErrInBounds", 0x1007, _S_OUT_) // fmt.Errorf(_S_OUT_)
	}
	return nil
}

//!+

/**
 * Pop remueve un nodo en la posición i. Base 0
 * @param {...T}
 * @return {T}
 * @return {error}
 */
func (list *List) Pop(i ...T) (T, error) {
	i_dbl := core.ArgOptional(list.size-1, i).(int)
	node := list.Head
	err := list._withinRange(i_dbl)
	if err == nil {
		node = list.Locate(i_dbl)
		remove := node.Next()
		node.next = remove.Next()
		list.size--
		Events.Emit("Pop", list, remove)
		return remove.Value(), err
	} else {
		return core.Literals().NullString, err
	}
}

//!-
/**
 * Remove la lista entera
 */
func (list *List) Remove() {
	*list = *__set()
	Events.Emit("Remove")
}

/**
 * Nos dice si contiene v y el número del registro. Base 0
 * @param {T}
 * @return {bool}
 * @return {int}
 */
func (list *List) Contains(v T) (bool, int) {
	if w := list.Index(v); w >= 0 {
		return true, w
	}
	return false, -1
}

/**
 * PushList inserta una copia de otra List
 * @param {*List}
 */
func (list *List) PushList(other *List) {
	for node, e := other.Head.next, 0; node != nil; node, e = node.Next(), e+1 {
		list.Insert(e, Paragraph{node.Value(), node.Info()})
	}
}

/**
 * Mueve un elemento a otra posición (pos Base 0, t Base 0)
 * @param {int}
 * @param {T}
 */
func (list *List) MoveTo(
	pos int,
	t T) {
	if v := list.Index(t); v >= 0 {
		list.Pop(v)
		list.Insert(pos, t)
		Events.Emit("Move", list, v)
	}
}

