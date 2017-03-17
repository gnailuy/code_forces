package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	face_mapping := map[string]int{
		"Tetrahedron":  4,
		"Cube":         6,
		"Octahedron":   8,
		"Dodecahedron": 12,
		"Icosahedron":  20}
	faces := 0

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err != io.EOF {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			faces += face_mapping[s]
		}
	}

	fmt.Println(faces)
}
