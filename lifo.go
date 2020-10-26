package lifo

// Lifo type
type Lifo struct {
	v []string
}

// New Lifo
func New() Lifo {
	return Lifo{}
}

// Add some data into the queue
func Add(l *Lifo, value string) {
	l.v = append(l.v, value)
}

// Remove item from queue
func Remove(l *Lifo) string {
	if IsEmpty(l) {
		return ""
	}
	i := len(l.v) - 1
	value := l.v[i]
	l.v = l.v[:i]
	return value
}

// Check if queue is empty
func IsEmpty(l *Lifo) bool {
	if len(l.v) == 0 {
		return true
	}
	return false
}
