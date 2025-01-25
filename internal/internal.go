package fileTree

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileNode struct {
	name     string
	isDir    bool
	children []*FileNode
}

func (n *FileNode) AddChild(child *FileNode) {
	n.children = append(n.children, child)
}

// TODO: Fix this logic. needs to be more dynamic. Current logic only works in very specific file structure
func (n *FileNode) PrintTree(spacer string, level int, isLast bool, parentIsLast bool) {
	prefix := ""
	suffix := ""

	switch level {
	case 0:
		if n.isDir {
			suffix = "/"
		}

	case 1:
		spacer = spacer + "├─ "
	default:
		spacer = "│  "
		for i := 0; i < level-2; i++ {
			if !parentIsLast {
				spacer += "│  "
			} else {
				spacer += "   "
			}
		}
		if isLast {
			prefix = "└─ "
		} else {
			prefix = "├─ "
		}
	}

	if n.isDir {
		suffix = "/"
	}

	fmt.Printf("%s%s%s%s\n", spacer, prefix, n.name, suffix)

	for i, child := range n.children {
		child.PrintTree(spacer, level+1, i == len(n.children)-1, isLast)
	}
}

func BuildFileTree(rootPath string) (*FileNode, error) {
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}

	root := &FileNode{
		name:  info.Name(),
		isDir: info.IsDir(),
	}

	if info.IsDir() {
		pathFiles, err := os.ReadDir(rootPath)
		if err != nil {
			return nil, err
		}

		for _, file := range pathFiles {
			childPath := filepath.Join(rootPath, file.Name())
			childNode, err := BuildFileTree(childPath)
			if err != nil {
				return nil, err
			}

			root.AddChild(childNode)
		}
	}

	return root, nil
}
