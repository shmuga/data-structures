package data_structures

type Traversable interface {
	Traverse(IVisitor)
}
