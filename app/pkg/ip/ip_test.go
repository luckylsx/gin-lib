package ip

import (
	"fmt"
	"testing"
)

type ips struct {
	ipString string
	ipInt    int
}

var tests = []ips{
	{
		ipString: "192.168.0.1",
		ipInt:    3232235521,
	},
	{
		ipString: "192.168.10.10",
		ipInt:    3232238090,
	},
	{
		ipString: "122.101.10.9",
		ipInt:    2053442057,
	},
	{
		ipString: "122.101.10.8",
		ipInt:    2053442056,
	},
}

func TestToInt(t *testing.T) {
	fmt.Println("ToInt test:")
	for _, v := range tests {
		if actual, _ := ToInt(v.ipString); v.ipInt != actual {
			t.Errorf("ip to int failed! ip is : %q, expected int ip is : %d, but got int ip is %d\n",
				v.ipString, v.ipInt, actual)
		}
	}
}

func TestToString(t *testing.T) {
	fmt.Println("ToString test:")
	for _, v := range tests {
		if actual, _ := ToString(v.ipInt); v.ipString != actual {
			t.Errorf("ip to string failed! int ip is : %d, expected string ip is : %q, but got int ip is %q\n",
				v.ipInt, v.ipString, actual)
		}
	}
}
