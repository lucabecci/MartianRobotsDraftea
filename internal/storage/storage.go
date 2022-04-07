package storage

// var of the storage to use
var storage LinkedList

// Node represents a node of linked list
type Node struct {
	value map[string]interface{}
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// Insert is responsible for inserting the interface received in a new node
// It is an impure function and does not return data.
func Insert(val map[string]interface{}) {
	n := Node{}
	n.value = val
	if storage.len == 0 {
		storage.head = &n
		storage.len++
		return
	}
	ptr := storage.head
	for i := 0; i < storage.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			storage.len++
			return
		}
		ptr = ptr.next
	}
}

// GetByAxis obtains and validates that there is a robot with that x and y in the storage
// Returns true or false depending on the validation
func GetByAxis(x, y int) (map[string]interface{}, bool) {
	x_str := float64(x)
	y_str := float64(y)
	if storage.len == 0 {
		return nil, false
	}
	ptr := storage.head
	for i := 0; i < storage.len; i++ {
		if ptr.value["X"] == x_str && ptr.value["Y"] == y_str {
			return ptr.value, true
		} else {
			ptr = ptr.next
		}
	}
	return nil, false
}
