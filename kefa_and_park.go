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

	edges := make([][]int, n)
	const init_size int = 100
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

			if edges[x-1] == nil {
				edges[x-1] = make([]int, 0, init_size)
				edges[x-1] = append(edges[x-1], 0)
			}
			edges[x-1][0] += 1
			edges[x-1] = append(edges[x-1], y)

			if edges[y-1] == nil {
				edges[y-1] = make([]int, 0, init_size)
				edges[y-1] = append(edges[y-1], 0)
			}
			edges[y-1][0] += 1
			edges[y-1] = append(edges[y-1], x)
		}
	}

	root := &Vertex{1, cats[0], nil, nil, nil}
	fmt.Println(build_sub_tree(root, edges, cats, root.cat, m))
}

func build_sub_tree(node *Vertex, edges [][]int, cats []int, cat, max_cat int) int {
	var last_child *Vertex = nil

	var cut bool = false
	for i := 1; i <= edges[node.idx-1][0]; i++ {
		var child_idx int = edges[node.idx-1][i]
		if node.parent == nil || child_idx != node.parent.idx {
			if cat+cats[child_idx-1] <= max_cat {
				new_child := &Vertex{child_idx, cats[child_idx-1], node, nil, nil}
				if node.children == nil {
					node.children = new_child
					last_child = new_child
				} else {
					last_child.next_sibling = new_child
					last_child = new_child
				}
			} else {
				cut = true
			}
		}
	}

	if node.children == nil {
		if cut {
			return 0
		} else {
			return 1
		}
	}

	res := 0
	for child := node.children; child != nil; child = child.next_sibling {
		var c int
		if child.cat == 0 {
			c = 0
		} else {
			c = cat + 1
		}
		res += build_sub_tree(child, edges, cats, c, max_cat)
	}
	return res
}
