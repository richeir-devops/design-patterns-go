package flyweight

import "fmt"

///////////////////////////////////////////
// 享元模式（Flyweight）:
// 核心是共享内存（对象），少创建新对象
///////////////////////////////////////////

///////////////////////////////////////////

type KarakTea struct {
	Description string
}

///////////////////////////////////////////

///////////////////////////////////////////

type TeaMaker struct {
	AvailableTea []*KarakTea
}

func (tm *TeaMaker) Make(desc string, table int) *KarakTea {
	newTea := &KarakTea{Description: desc}
	tm.AvailableTea[table] = newTea
	return newTea
}

///////////////////////////////////////////

///////////////////////////////////////////

type TeaShop struct {
	orders   []*KarakTea
	teaMaker *TeaMaker
}

func (ts *TeaShop) TakeOrder(teaType string, table int) {
	ts.orders[table] = ts.teaMaker.Make(teaType, table)
}

func (ts *TeaShop) Serve() {
	for index, order := range ts.orders {
		if order != nil {
			fmt.Printf("Serving tea to table# %d : %s \n", index, order.Description)
		}

	}
}

///////////////////////////////////////////
