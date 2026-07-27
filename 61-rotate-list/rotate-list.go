/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getLen(head *ListNode) int {
    cnt := 0
    for head != nil {
        cnt++
        head = head.Next
    }
    return cnt
}
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    fast, slow := head, head
    i := 0
    n := getLen(head)
    k = k % n
    for i < k {
        fast = fast.Next
        i++
    }  

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }

    fast.Next = head
    node := slow.Next
    slow.Next = nil

    return node
}