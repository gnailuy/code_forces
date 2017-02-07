package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Edge struct {
	x    int
	y    int
	next *Edge
}

type Vertex struct {
	idx          int
	cat          int
	parent       *Vertex
	children     *Vertex
	next_sibling *Vertex
}

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	cats := make([]int, n)
	_v, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	vertices := strings.Split(re.ReplaceAllString(_v, ""), " ")
	if len(vertices) != n {
		return
	}
	for i := 0; i < n; i++ {
		v, err := strconv.Atoi(vertices[i])
		if err != nil {
			return
		}
		cats[i] = v
	}

	var edges *Edge = &Edge{0, 0, nil}
	var p *Edge = edges
	for i := 0; i < n-1; i++ {
		if _e, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			e := strings.Split(re.ReplaceAllString(_e, ""), " ")
			if len(e) != 2 {
				return
			}
			x, err := strconv.Atoi(e[0])
			if err != nil {
				return
			}
			y, err := strconv.Atoi(e[1])
			if err != nil {
				return
			}
			if x > y {
				x, y = y, x
			}
			new_edge := &Edge{x, y, nil}
			p.next = new_edge
			p = p.next
		}
	}

	root := &Vertex{1, cats[0], nil, nil, nil}
	build_sub_tree(root, edges, cats)

	fmt.Println(visit_sub_tree(root, root.cat, m))
}

func visit_sub_tree(node *Vertex, current_cat, max_cat int) int {
	if node.children == nil {
		return 1
	} else {
		res := 0
		for child := node.children; child != nil; child = child.next_sibling {
			if child.cat == 0 {
				res += visit_sub_tree(child, 0, max_cat)
			} else if child.cat+current_cat <= max_cat {
				res += visit_sub_tree(child, child.cat+current_cat, max_cat)
			}
		}
		return res
	}
}

func build_sub_tree(node *Vertex, edges *Edge, cats []int) {
	var last_child *Vertex = nil

	var prev *Edge = edges
	var e *Edge = edges.next
	for e != nil {
		if e.x != node.idx && e.y != node.idx {
			prev = e
			e = e.next
			continue
		}

		var child_idx int
		if e.x == node.idx {
			child_idx = e.y
		} else if e.y == node.idx {
			child_idx = e.x
		}
		if node.parent == nil || child_idx != node.parent.idx {
			new_child := &Vertex{child_idx, cats[child_idx-1], node, nil, nil}
			if node.children == nil {
				node.children = new_child
				last_child = new_child
			} else {
				last_child.next_sibling = new_child
				last_child = new_child
			}
			prev.next = e.next
			e = e.next
		} else {
			prev = e
			e = e.next
		}
	}
	for child := node.children; child != nil; child = child.next_sibling {
		build_sub_tree(child, edges, cats)
	}
}
