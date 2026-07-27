/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getLen(head *ListNode) int {
    if head == nil {
        return 0
    }
    cnt := 0
    for head != nil {
        cnt++
        head = head.Next
    }
    return cnt
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    N := getLen(head)
    if N == n {
        return head.Next
    }
    target := N - n
    fmt.Println(N, n, target)
    var node *ListNode = head
    for node != nil && target > 1 {
        node = node.Next
        target--
    }
    node.Next = node.Next.Next
    fmt.Println(node.Val, target)
    return head
}