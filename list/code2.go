package list

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var l ListNode
    var flag int

    cur := &l
    cur1 := l1
    cur2 := l2

    for {

        if cur1 != nil {
            cur.Val += cur1.Val
        }

        if cur2 != nil {
            cur.Val += cur2.Val
        }

        cur.Val += flag
        if cur.Val >= 10 {
            cur.Val -= 10
            flag = 1
        } else {
            flag = 0
        }

        if cur1 != nil  {
            cur1 = cur1.Next
        }

        if cur2 != nil{
            cur2 = cur2.Next
        }

        if cur1 == nil && cur2 == nil {
            break
        }

        cur.Next = &ListNode{
            Val:  0,
            Next: nil,
        }

        cur = cur.Next
    }

    if flag == 1 {
        cur.Next = &ListNode{
            Val:  flag,
            Next: nil,
        }
    }

    return &l

}
