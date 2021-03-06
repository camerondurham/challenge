package main

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    if node.Next.Next == nil {
        node.Next = nil
    } else {
        deleteNode(node.Next)
    }
}
