/**
 * https://leetcode.com/problems/execute-asynchronous-functions-in-parallel
 * @param {Array<Function>} functions
 * @return {Promise<any>}
 */
var promiseAll = async function (functions) {
  return new Promise((resolve, reject) => {
    let count = functions.length,
      results = new Array(count);
    for (let i = 0; i < count; i++) {
      functions[i]()
        .then((val) => {
          results[i] = val;
          if (--count === 0) resolve(results);
        })
        .catch((reason) => reject(reason));
    }
  });
};

/**
 * const promise = promiseAll([() => new Promise(res => res(42))])
 * promise.then(console.log); // [42]
 */
module.exports = { promiseAll };
