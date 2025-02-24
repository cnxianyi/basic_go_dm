package method

// 方法
// 规定 必须实现这三个方法
type List interface {
	Add(v any, i int) error
	Append(v any) error
	Delete(i int) (v any, err error)
}

type LinkList struct {
	next *LinkList
}
