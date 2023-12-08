package main

import (
	"strings"
)

type TreeNode struct {
	key   string
	left  *TreeNode
	right *TreeNode
}

func part1(input string) any {
	splitted := strings.Split(input, "\n")
	dirs := splitted[0]

	nodes := make(map[string]*TreeNode)

	for _, line := range splitted[2:] {
		ckey, lkey, rkey := line[0:3], line[7:10], line[12:15]
		left := createNodeIfAbsent(nodes, lkey, nil, nil)
		right := createNodeIfAbsent(nodes, rkey, nil, nil)
		createNodeIfAbsent(nodes, ckey, left, right)
	}

	return traverse(nodes["AAA"], dirs, func(k string) bool { return k == "ZZZ" })
}

func part2(input string) any {
	splitted := strings.Split(input, "\n")
	dirs := splitted[0]

	nodes := make(map[string]*TreeNode)

	roots := make([]*TreeNode, 0)
	for _, line := range splitted[2:] {
		ckey, lkey, rkey := line[0:3], line[7:10], line[12:15]
		left := createNodeIfAbsent(nodes, lkey, nil, nil)
		right := createNodeIfAbsent(nodes, rkey, nil, nil)
		root := createNodeIfAbsent(nodes, ckey, left, right)
		if ckey[len(ckey)-1] == 'A' {
			roots = append(roots, root)
		}
	}

	res := 1
	for _, root := range roots {
		res = lcm(res, traverse(root, dirs, func(k string) bool { return k[len(k)-1] == 'Z' }))
	}

	return res
}

func createNodeIfAbsent(nodes map[string]*TreeNode, key string, left, right *TreeNode) *TreeNode {
	if nodes[key] == nil {
		nodes[key] = &TreeNode{
			key:   key,
			left:  left,
			right: right,
		}
	} else {
		if left != nil {
			nodes[key].left = left
		}
		if right != nil {
			nodes[key].right = right
		}
	}

	return nodes[key]
}

func traverse(root *TreeNode, dirs string, isFinish func(k string) bool) int {
	cur := root
	curI := 0
	cnt := 0

	for cur != nil && !isFinish(cur.key) {
		cnt++
		if dirs[curI] == 'L' {
			cur = cur.left
		} else {
			cur = cur.right
		}

		curI = (curI + 1) % len(dirs)
	}

	return cnt
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
