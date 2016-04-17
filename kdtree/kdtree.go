package kdtree

import (
	"fmt"
	"strings"
)

type kdnode struct {
	item        interface{}
	values      []interface{}
	dim, disc   int
	left, right *kdnode
}

func (n *kdnode) getvalues() []interface{} {
	return n.values
}

func (n *kdnode) getvalue(disc int) interface{} {
	if disc < 0 || disc >= len(n.values) {
		return nil
	}
	return n.values[disc]
}

func (n *kdnode) setvals(values []interface{}) {
	if len(values) == len(n.values) {
		n.values = values
	}
}

type KDTree struct {
	root *kdnode
	dim  int
}

func (k *KDTree) Insert(item interface{}, values []interface{}) {
	if len(values) == k.dim {
		vals := make([]interface{}, len(values))
		copy(vals, values)
		var newnode = &kdnode{item, vals, k.dim, 0, nil, nil}
		if k.root == nil {
			k.root = newnode
		} else {
			k.insert(k.root, newnode)
		}
	}
}

func (k *KDTree) insert(subroot, newnode *kdnode) {
	var goright bool
	switch subroot.getvalue(subroot.disc).(type) {
	case int:
		goright = newnode.getvalue(subroot.disc).(int) >= subroot.getvalue(subroot.disc).(int)
	case float64:
		goright = newnode.getvalue(subroot.disc).(float64) >= subroot.getvalue(subroot.disc).(float64)
	}
	if goright {
		if subroot.right == nil {
			subroot.right = newnode
			newnode.disc = (subroot.disc + 1) % k.dim
		} else {
			k.insert(subroot.right, newnode)
		}
	} else {
		if subroot.left == nil {
			subroot.left = newnode
			newnode.disc = (subroot.disc + 1) % k.dim
		} else {
			k.insert(subroot.left, newnode)
		}
	}
}

func (k *KDTree) Traverse(style string) []string {
	var tree = k.traverse(k.root, style)
	return strings.FieldsFunc(tree, func(r rune) bool {
		return r == ' ' || r == '\n'
	})
}

func (k *KDTree) traverse(subroot *kdnode, style string) string {
	var tree string
	if subroot == nil {
		return tree
	}
	if style == "pre" {
		tree += fmt.Sprint(subroot.item) + " "
	}
	tree += k.traverse(subroot.left, style)
	if style == "in" {
		tree += fmt.Sprint(subroot.item) + " "
	}
	tree += k.traverse(subroot.right, style)
	if style == "post" {
		tree += fmt.Sprint(subroot.item) + " "
	}
	return tree
}
