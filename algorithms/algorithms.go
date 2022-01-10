package algorithms

type Algorithm string

const (
    BinaryTree Algorithm = "binary-tree"
    Sidewinder Algorithm = "sidewinder"
)

func Algorithms() []string {
    return []string{BinaryTree.String(), Sidewinder.String()}
}

func (a Algorithm) String() string {
    return string(a)
}
