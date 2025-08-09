package avl

import (
	"cmp"
)

type AVLNode[T cmp.Ordered] struct {
	value  T
	left   *AVLNode[T]
	right  *AVLNode[T]
	size int
}

type AVLTree[T cmp.Ordered] struct {
	root *AVLNode[T]
}

// NewAVLTree crea un nuevo árbol AVL vacío
func NewAVLTree[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

// height devuelve la altura de un nodo (maneja nodos nulos)
func (n *AVLNode[T]) height() int {
	if n == nil {
		return 0
	}
	return n.size
}

// max devuelve el máximo entre dos enteros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// updateHeight actualiza la altura de un nodo basado en sus hijos
func (n *AVLNode[T]) updateHeight() {
	n.size = 1 + max(n.left.height(), n.right.height())
}

// balanceFactor calcula el factor de balance de un nodo
func (n *AVLNode[T]) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.height() - n.right.height()
}

// rotateRight realiza una rotación a la derecha
func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// rotateLeft realiza una rotación a la izquierda
func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// balance rebalancea el árbol AVL después de una inserción o eliminación
func (n *AVLNode[T]) balance() *AVLNode[T] {
	n.updateHeight()

	balance := n.balanceFactor()

	// Caso izquierda-izquierda
	if balance > 1 && n.left.balanceFactor() >= 0 {
		return n.rotateRight()
	}

	// Caso izquierda-derecha
	if balance > 1 && n.left.balanceFactor() < 0 {
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	// Caso derecha-derecha
	if balance < -1 && n.right.balanceFactor() <= 0 {
		return n.rotateLeft()
	}

	// Caso derecha-izquierda
	if balance < -1 && n.right.balanceFactor() > 0 {
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

// Insert añade un nuevo valor al árbol AVL
func (t *AVLTree[T]) Insert(value T) {
	t.root = t.insert(t.root, value)
}

func (t *AVLTree[T]) insert(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return &AVLNode[T]{value: value, size: 1}
	}

	if value < node.value {
		node.left = t.insert(node.left, value)
	} else if value > node.value {
		node.right = t.insert(node.right, value)
	} else {
		// Valor duplicado, no se hace nada
		return node
	}

	return node.balance()
}

// Delete elimina un valor del árbol AVL
func (t *AVLTree[T]) Delete(value T) {
	t.root = t.delete(t.root, value)
}

func (t *AVLTree[T]) delete(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = t.delete(node.left, value)
	} else if value > node.value {
		node.right = t.delete(node.right, value)
	} else {
		// Nodo con un solo hijo o sin hijos
		if node.left == nil || node.right == nil {
			var temp *AVLNode[T]
			if node.left != nil {
				temp = node.left
			} else {
				temp = node.right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
		} else {
			// Nodo con dos hijos: obtener el sucesor inorden (mínimo en el subárbol derecho)
			temp := t.minValueNode(node.right)
			node.value = temp.value
			node.right = t.delete(node.right, temp.value)
		}
	}

	if node == nil {
		return node
	}

	return node.balance()
}

// minValueNode encuentra el nodo con el valor mínimo en un árbol
func (t *AVLTree[T]) minValueNode(node *AVLNode[T]) *AVLNode[T] {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// Search busca un valor en el árbol AVL
func (t *AVLTree[T]) Search(value T) bool {
	return t.search(t.root, value)
}

func (t *AVLTree[T]) search(node *AVLNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value < node.value {
		return t.search(node.left, value)
	} else if value > node.value {
		return t.search(node.right, value)
	}

	return true
}

// InOrder recorre el árbol en orden (izquierda, raíz, derecha)
func (t *AVLTree[T]) InOrder() []T {
	var result []T
	t.inOrder(t.root, &result)
	return result
}

func (t *AVLTree[T]) inOrder(node *AVLNode[T], result *[]T) {
	if node != nil {
		t.inOrder(node.left, result)
		*result = append(*result, node.value)
		t.inOrder(node.right, result)
	}
}