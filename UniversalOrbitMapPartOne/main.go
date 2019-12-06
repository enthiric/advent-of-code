package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tree struct {
	Root *Node
}

func (t *Tree) CountOrbitsFrom(n *Node) int {
	if n == nil {
		return 0
	}

	return t.CountOrbits(n, 0)
}

func (t *Tree) CountOrbits(n *Node, c int) int {
	if n.Parent != nil {
		x := n
		for {
			if x.Parent == nil {
				break
			}

			x = x.Parent
			c++
		}
	}

	for _, node := range n.Childs {
		c = t.CountOrbits(node, c)
	}

	return c
}

type Node struct {
	Name   string
	Parent *Node
	Childs []*Node
}

func (n *Node) FindOrbit(orbit string) *Node {
	for _, node := range n.Childs {
		if node == nil {
			continue
		}

		if node.Name == orbit {
			return node
		}

		if found := node.FindOrbit(orbit); found != nil {
			return found
		}
	}

	return nil
}

func NewNode(orbit string, parent *Node) *Node {
	return &Node{Name: orbit, Parent: parent}
}

type List []string

func (l List) Find(orbit string) (string, int) {
	for k := range l {
		if strings.HasPrefix(l[k], orbit) {
			return l[k], k
		}
	}

	return "", -1
}

func (l List) Remove(idx int) List {
	if len(l) == 1 {
		return List{}
	}

	return append(l[:idx], l[idx+1:]...)
}

func main() {
	path, err := filepath.Abs("./UniversalOrbitMapPartOne/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	list := List{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	path, idx := list.Find("COM")
	orbits := strings.Split(path, ")")
	list = list.Remove(idx)

	tree := Tree{}
	tree.Root = NewNode(orbits[0], nil)
	tree.Root.Childs = append(tree.Root.Childs, NewNode(orbits[1], tree.Root))

	for {
		if len(list) == 0 {
			break
		}

		for i := len(list) - 1; i >= 0; i-- {
			orbits := strings.Split(string(list[i]), ")")
			if len(orbits) == 2 {
				node := tree.Root.FindOrbit(orbits[0])
				if node == nil {
					continue
				}

				list = list.Remove(i)
				node.Childs = append(node.Childs, NewNode(orbits[1], node))
			}
		}
	}

	log.Print(tree.CountOrbitsFrom(tree.Root))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
