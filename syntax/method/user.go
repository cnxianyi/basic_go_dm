package method

type User interface {
	ChangeName(*Users) error
	ChangeAge(*Users) error
}

type Users struct {
	Name string
	Age  int
}

// 类型别名. 与原类型相等
type UserRemark = Users

// 直接传值使用的是 值传递. 无法修改原本的结构体值
// 需要传递指针. 即使是值传递. 但是值传递的指针仍指向原结构体
func (u Users) ChangeAge(age int) error {
	// unused write to field Age
	u.Age = age
	return nil
}

// 传递指针
func (u *Users) ChangeName(name string) error {
	u.Name = name
	return nil
}
