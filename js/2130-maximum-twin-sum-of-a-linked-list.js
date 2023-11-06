/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {number}
 */
var pairSum = function (head) {
  var slow = (fast = head);
  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next.next;
  }

  var second = slow,
    prev = null;
  while (second) {
    var tmp = second.next;
    second.next = prev;
    prev = second;
    second = tmp;
  }

  var maxSum = 0;
  while (prev) {
    maxSum = Math.max(maxSum, head.val + prev.val);
    head = head.next;
    prev = prev.next;
  }

  return maxSum;
};

module.exports = { pairSum };
