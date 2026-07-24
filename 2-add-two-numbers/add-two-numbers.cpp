/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* p1 = l1;
        ListNode* p2 = l2;
        ListNode* dummy = new ListNode(0);
        ListNode* ans = dummy;
        int carry = 0;

        while(p1 != NULL || p2 != NULL) {
            int x = (p1 ? p1->val : 0);
            int y = (p2 ? p2->val : 0);

            int sum = x + y + carry;
            carry = sum / 10;

            ans->next = new ListNode(sum % 10);
            ans = ans->next;

            if (p1) p1 = p1->next;
            if (p2) p2 = p2->next;
        }

        if(carry) {
            ans->next = new ListNode(carry);
        }

        return dummy->next;

    }
};