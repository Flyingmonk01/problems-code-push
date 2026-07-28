/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return head
    }

    l_ptr := head
    r_ptr := head

    dummy := &ListNode{Next: head}
    prev := dummy

    var r_next *ListNode
    i := 1
    for i < left {
        i++
        prev = l_ptr
        l_ptr = l_ptr.Next
    }
    i = 1
    for i < right {
        i++
        r_ptr = r_ptr.Next
    }
    if r_ptr.Next != nil {
        r_next = r_ptr.Next
    }   

    prev.Next = nil
    r_ptr.Next = nil

    var nxt *ListNode
    leftNode := l_ptr

    for l_ptr != nil {
        curr_nxt := l_ptr.Next
        l_ptr.Next = nxt
        nxt = l_ptr
        l_ptr = curr_nxt
    }

    prev.Next = nxt
    leftNode.Next = r_next
    return dummy.Next

}