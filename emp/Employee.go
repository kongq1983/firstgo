package emp

import "fmt"

type Employee struct {
	Id    int
	Username string
	Name  string
	Age int
}


func (s *Employee) ToString() string {
	return fmt.Sprintf("[id => %d name => %v, age => %v]", s.Id, s.Name, s.Age)
}


func (s *Employee) IsEqualAge(e2 *Employee) bool {

	if s.Age==e2.Age {
		return true
	}

	return false

}

func (s *Employee) IsEqualAge1(e2 Employee) bool {

	if s.Age==e2.Age {
		return true
	}

	return false

}


func main() {

	var e1 = Employee{Id:1,Username: "admin",Name: "管理员",Age: 20}
	var e2 = Employee{Id:2,Username: "guest",Name: "匿名用户",Age: 20}
	var e3 = Employee{Id:3,Username: "king",Name: "管理员",Age: 22}

	fmt.Println("e1",e1.ToString())
	fmt.Println("e2",e2.ToString())
	fmt.Println("e3",e3.ToString())

	var equal = e1.IsEqualAge(&e2)
	fmt.Println("e1.age==e2.age",equal)

	var equal1 = e1.IsEqualAge1(e3)
	fmt.Println("e1.age==e3.age",equal1)


}
