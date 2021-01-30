package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// String formats a DMS coordinate.
func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

type location struct {
	Name string     `JSON:"name"` //记得大写
	Lat  coordinate `JSON:"latitude"`
	Long coordinate `JSON:"longtitude"`
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct { //用Marshal序列化，Json的key键值是个struct结构
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(), //与key对应的value值
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

func main() {
	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}

	//因为coordinate满足了MarshalJSON() ([]byte, error)隐式接口，所以可以直接被json.Marshal调用
	message, err := json.MarshalIndent(elysium, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(message)) //message是[]byte类型，转化成string类型便于查看

}
