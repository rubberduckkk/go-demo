package wepie

import "testing"

func TestLastNthNode(t *testing.T) {
	type data struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2
	node2.Next = node3

	testData := []data{
		{
			name: "case1",
			head: node1,
			n:    2,
			want: node2,
		},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			res := deleteLastNthNode(d.head, d.n)
			if res != d.want {
				t.Fatalf("test failed, want=%+v, get=%+v\n", d.want, res)
			}
		})
	}
}
