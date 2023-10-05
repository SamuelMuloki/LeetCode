// Definition for singly-linked list.
function ListNode(val, next) {
  var obj = {};
  obj.val = val === undefined ? 0 : val;
  obj.next = next === undefined ? null : next;

  return obj;
}

module.exports = { ListNode };
