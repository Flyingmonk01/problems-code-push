/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    // length := 0
    // var node *ListNode = head
    // for node != nil {
    //     length++
    //     node = node.Next
    // }
    // target := length / 2 - 1
    // var new_node *ListNode = head
    // for target != 0 {
    //     new_node = new_node.Next
    //     target--
    // }
    // new_node.Next = new_node.Next.Next
    slow, fast := head, head.Next.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    slow.Next = slow.Next.Next
    return head
    
}