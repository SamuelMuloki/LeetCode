/**
 * @param {number[][]} rooms
 * @return {boolean}
 */
var canVisitAllRooms = function (rooms) {
  var seen = new Array(rooms.length).fill(false);
  seen[0] = true;
  var st = [0];
  while (st.length > 0) {
    var node = st.pop();
    rooms[node].forEach((key) => {
      if (!seen[key]) {
        seen[key] = true;
        st.push(key);
      }
    });
  }

  return seen.every((val) => val == true);
};

module.exports = { canVisitAllRooms };
