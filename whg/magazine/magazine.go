package magazine

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address // 匿名struct字段 == 嵌入外部struct，可使用Address访问或者可直接访问Street、City等
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address // 非匿名struct字段，必须使用HomeAddress名称访问
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
