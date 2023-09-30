package main

import (
//    "os"
    "strconv"
    "time"
    "sort"
	"fmt"
	"math/rand"
    "runtime"
)

func genData(seed int64, size int, maxNum int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Intn(maxNum)
	}

	rand.Seed(seed) 
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

type AVLTree struct {
	Root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rotateRight(y *Node) *Node {
	x := y.Left
	temp := x.Right

	x.Right = y
	y.Left = temp

	updateHeight(y)
	updateHeight(x)

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	temp := y.Left

	y.Left = x
	x.Right = temp

	updateHeight(x)
	updateHeight(y)

	return y
}


func (t *AVLTree) Insert(key int) {
	t.Root = t.insertRecursive(t.Root, key)
}

func (t *AVLTree) insertRecursive(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Height: 1}
	}

	if key < root.Key {
		root.Left = t.insertRecursive(root.Left, key)
	} else if key > root.Key {
		root.Right = t.insertRecursive(root.Right, key)
	} else {
		return root 
	}

	updateHeight(root)

	balance := balanceFactor(root)

	if balance > 1 {
		if key < root.Left.Key {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}

	if balance < -1 {
		if key > root.Right.Key {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}

func (t *AVLTree) Find(key int) *Node {
	return t.findRecursive(t.Root, key)
}

func (t *AVLTree) findRecursive(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key == root.Key {
		return root
	} else if key < root.Key {
		return t.findRecursive(root.Left, key)
	} else {
		return t.findRecursive(root.Right, key)
	}
}

func (t *AVLTree) Delete(key int) {
	t.Root = t.deleteRecursive(t.Root, key)
}

func (t *AVLTree) deleteRecursive(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Key {
		root.Left = t.deleteRecursive(root.Left, key)
	} else if key > root.Key {
		root.Right = t.deleteRecursive(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *Node
			if root.Left != nil {
				temp = root.Left
			} else {
				temp = root.Right
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}

			temp = nil
		} else {
			temp := findMinNode(root.Right)
			root.Key = temp.Key
			root.Right = t.deleteRecursive(root.Right, temp.Key)
		}
	}

	if root == nil {
		return root
	}

	updateHeight(root)

	balance := balanceFactor(root)

	if balance > 1 {
		if balanceFactor(root.Left) >= 0 {
			return rotateRight(root)
		} else {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	}

	if balance < -1 {
		if balanceFactor(root.Right) <= 0 {
			return rotateLeft(root)
		} else {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}

	return root
}


func findMinNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func printTree(node *Node) {
	if node == nil {
		return
	}

	printTree(node.Left)
	fmt.Printf("%d ", node.Key)
	printTree(node.Right)
}

func benchmark(seed int64, size, maxNum, order int) {
	var m runtime.MemStats
    runtime.ReadMemStats(&m)
    allocMemStart := m.TotalAlloc
    tree := NewAVLTree()
	keys := genData(seed, size, maxNum)

    if order == 1 {
        sort.Slice(keys, func(i, j int) bool {
            return keys[i] < keys[j]
        })
    } else if order == -1 {
        sort.Slice(keys, func(i, j int) bool {
            return keys[i] >= keys[j]
        })
    }

	for _, key := range keys {
		tree.Insert(key)
	}

    //if (os.Args[1] == "insert") {
        start := time.Now()
        tree.Insert(-1)
        tree.Insert(maxNum/2)
        tree.Insert(maxNum)
        fmt.Printf("%d,%s,%d\n", size, "Insert", time.Since(start))
    //}

    //if (os.Args[1] == "find") {
        start = time.Now()
        tree.Find(-1)
        tree.Find(maxNum/2)
        tree.Find(maxNum)
        fmt.Printf("%d,%s,%d\n", size, "Find", time.Since(start))
    //}    

    //if (os.Args[1] == "delete") {
        start = time.Now()
        tree.Delete(-1)
        tree.Delete(maxNum/2)
        tree.Delete(maxNum)
        fmt.Printf("%d,%s,%d\n", size, "Delete", time.Since(start))
    //}

    runtime.ReadMemStats(&m)
    fmt.Printf("%d,%s,%d\n", size, "Memory", m.TotalAlloc-allocMemStart)
}


func toInt(str string) int {
    val, _ := strconv.Atoi(str)
    return val
}


func main() {
    fmt.Println("# of elements,operation,duration (ns)")
    j:=100
    for i:=1000; i<10000; i+=j {
        benchmark( 1, i, i, 0 )
        if i < 10000 {
            j = 100
        } else if i < 100000 {
            j = 1000
        } else if i < 1000000 {
            j = 10000
        }  else if i < 10000000 {
            j = 100000
        } 
    }
}

