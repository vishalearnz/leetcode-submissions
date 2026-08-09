/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    temp := head
    count := 0; 
    for i := 0 ; i < k ; i++ {
        if temp == nil {
            return head // Less than k nodes remain, return unreversed
        }
        temp = temp.Next
        count++
    }
    if count < k {
        return head
    }

    newHead := reverseKnodes(head, temp)
    head.Next =  reverseKGroup(temp, k)
    return newHead
    
}

func reverseKnodes(head *ListNode, temp *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head

    for curr != temp {
        next := curr.Next // 1. Save next node
        curr.Next = prev  // 2. Reverse current pointer
        prev = curr       // 3. Move prev forward
        curr = next       // 4. Move curr forward
    }

    return prev // 'prev' is now the NEW HEAD of this reversed sublist!
}