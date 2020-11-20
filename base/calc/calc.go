package calc

const a int = 100
const B int = 200

type stu2 struct {
	name string
}
type Stu struct {
	name string
	age  int8
}

func add(a, b int) int {
	return a + b
}

func Add(a, b int) (ret int) {
	ret = a + b
	return
}
