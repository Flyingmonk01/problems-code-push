/*
class Node {
public:
    int data;
    Node* next;
    Node* bottom;

    Node(int x) {
        data = x;
        next = NULL;
        bottom = NULL;
    }
};
*/

class Solution {
  public:
  
    Node* merge(Node* left, Node* right){
        if(!left) return right;
        if(!right) return left;
        
        Node* a = 0;
        if(left->data < right->data){
            a = left;
            left->bottom = merge(left->bottom, right);
        }
        else{
            a = right;
            right->bottom = merge(left, right->bottom);
        }
        return a;
    }
    Node *flatten(Node *head) {
        // code here
        if (!head) return nullptr;
        Node* head2 = flatten(head->next);
        return merge(head, head2);
    }
};