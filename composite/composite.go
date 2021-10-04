package composite

///////////////////////////////////////////
// 组合模式（Composite）:
// 各种对象都实现一个基础接口，这个接口定义了基础的方法；
// 有一个类似总管的类，使用基础接口把各个对象聚合到一起，并实施各种Admin级别的骚操作
///////////////////////////////////////////

///////////////////////////////////////////

type Employee interface {
	Constructor(string, float32)
	GetSalary() float32
}

///////////////////////////////////////////

///////////////////////////////////////////

type Developer struct {
	name   string
	salary float32
}

func (d *Developer) Constructor(name string, salary float32) {
	d.name = name
	d.salary = salary
}

func (d *Developer) GetSalary() float32 {
	return d.salary
}

///////////////////////////////////////////

///////////////////////////////////////////

type Designer struct {
	name   string
	salary float32
}

func (d *Designer) Constructor(name string, salary float32) {
	d.name = name
	d.salary = salary
}

func (d *Designer) GetSalary() float32 {
	return d.salary
}

///////////////////////////////////////////

///////////////////////////////////////////

type Organization struct {
	emps []Employee
}

func (o *Organization) AddEmployee(e Employee) {
	o.emps = append(o.emps, e)
}

func (o *Organization) GetNetSalaries() float32 {
	result := 0

	for _, e := range o.emps {
		result += int(e.GetSalary())
	}

	return float32(result)
}

///////////////////////////////////////////
