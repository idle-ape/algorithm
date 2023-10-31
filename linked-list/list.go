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
