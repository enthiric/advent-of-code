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

func (t *Tree) CountOrbitsFrom(n *Node, to string) int {
	if n == nil {
		return 0
	}

	return t.CountOrbits(n, to, 0)
}

func (t *Tree) CountOrbits(n *Node, to string, c int) int {
	var y *Node
	x := n
	f := false
	for {
		if x.Parent == nil {
			log.Fatal("we failed")
		}

		if !f {
			if found := x.FindOrbit(to); found != nil {
				// keep the cross orbit
				y = x
				f = true
				x = found.Parent
			}
		}

		x = x.Parent

		if f {
			if x.Name == y.Name {
				break
			}
		}

		c++
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
	path, err := filepath.Abs("./PartTwo/input.txt")
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

	path, idx := list.Find("COM)")
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

	log.Print(tree.CountOrbitsFrom(tree.Root.FindOrbit("YOU"), "SAN"))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
