/**
 * @param {Function} fn
 * @return {Array}
 */
Array.prototype.groupBy = function (fn) {
  return this.reduce((output, element) => {
    var res = fn(element);

    if (res in output) {
      output[res].push(element);
    } else {
      output[res] = [element];
    }

    return output;
  }, {});
};

/**
 * [1,2,3].groupBy(String) // {"1":[1],"2":[2],"3":[3]}
 */
