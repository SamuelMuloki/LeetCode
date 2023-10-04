/**
 * @param {number[]} nums
 * @return {boolean}
 */
var find132pattern = function (nums) {
  var st = [],
    third = Number.MIN_SAFE_INTEGER;
  for (var i = nums.length - 1; i >= 0; i--) {
    if (nums[i] < third) {
      return true;
    }

    while (st.length > 0 && st[st.length - 1] < nums[i]) {
      third = st[st.length - 1];
      st.pop();
    }

    st.push(nums[i]);
  }

  return false;
};

module.exports = { find132pattern };
