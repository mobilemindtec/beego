package web

// Finally guaranteed execution if panic
type Finally interface {
	Finally()
}

// Recover executes in case of panic
type Recover interface {
	Recover(interface{})
}

// PreRender Execute on pre-render
type PreRender interface {
	PreRender(interface{})
}
