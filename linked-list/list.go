package linkedlist

import "fmt"

var (
	List1 = &List{
		Val: 1,
		Next: &List{
			Val: 2,
			Next: &List{
				Val: 4,
				Next: &List{
					Val: 6,
				},
			},
		},
	}
	List2 = &List{
		Val: 1,
		Next: &List{
			Val: 3,
			Next: &List{
				Val: 5,
				Next: &List{
					Val: 9,
				},
			},
		},
	}
	List3 = &List{
		Val: 1,
		Next: &List{
			Val: 4,
			Next: &List{
				Val: 3,
				Next: &List{
					Val: 2,
					Next: &List{
						Val: 5,
						Next: &List{
							Val: 2,
						},
					},
				},
			},
		},
	}
	node1 = &List{
		Val: 2,
	}
	node2 = &List{
		Val: 5,
	}
	node3 = &List{
		Val: 2,
	}
	node4 = &List{
		Val: 4,
	}
	node5 = &List{
		Val: 3,
	}
	CycleList = &List{
		Val: 1,
	}
)

type List struct {
	Val  int
	Next *List
}

func (l *List) String() string {
	str := ""
	if l == nil {
		return str
	}

	for p := l; p != nil; p = p.Next {
		if str == "" {
			str = fmt.Sprintf("%d", p.Val)
			continue
		}
		str = fmt.Sprintf("%s - %d", str, p.Val)
	}

	return str
}
