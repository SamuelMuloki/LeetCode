var { containsDuplicate } = require("./0217-contains-duplicate");
var { isAnagram } = require("./0242-valid-anagram");
var { convertToTitle } = require("./0168-excel-sheet-column-title");
var { createHelloWorld } = require("./2667-create-hello-world-function");
var { createCounter } = require("./2620-counter");
const { expect } = require("./2704-to-be-or-not-to-be");
const { createCounter2 } = require("./2665-counter-ii");
const { map } = require("./2635-apply-transform-over-each-element-in-array");
const { filter } = require("./2634-filter-elements-from-array");
const { reduce } = require("./2626-array-reduce-transformation");
const { compose } = require("./2629-function-composition");
const { once } = require("./2666-allow-one-function-call");
const { memoize } = require("./2623-memoize");
const { addTwoPromises } = require("./2723-add-two-promises");
const { cancellable } = require("./2715-timeout-cancellation");
const { timeLimit } = require("./2637-promise-time-limit");
const { TimeLimitedCache } = require("./2622-cache-with-time-limit");
const { debounce } = require("./2627-debounce");
const {
  promiseAll,
} = require("./2721-execute-asynchronous-functions-in-parallel");
const { isEmpty } = require("./2727-is-object-empty");
const { chunk } = require("./2677-chunk-array");
const { sortBy } = require("./2724-sort-by");
const { join } = require("./2722-join-two-arrays-by-id");

console.log(containsDuplicate([1, 2, 3, 1]));
console.log(isAnagram("rat", "car"));
console.log(convertToTitle(701));
console.log(createHelloWorld()());
const counter = createCounter(10);
console.log(counter());
console.log(expect(5).toBe(5));
const counter2 = createCounter2(5);
console.log(counter2.increment());
console.log(counter2.reset());
console.log(counter2.decrement());
console.log(
  map([1, 2, 3], function plusI(n, i) {
    return n + i;
  })
);
console.log(
  filter([-2, -1, 0, 1, 2], function plusOne(n) {
    return n + 1;
  })
);
console.log(
  reduce(
    [],
    function sum(accum, curr) {
      return 0;
    },
    25
  )
);
const composeFn = compose([]);
console.log(composeFn(42));

let ofn = (a, b, c) => a + b + c;
let onceFn = once(ofn);
console.log(onceFn(1, 2, 3));
console.log(onceFn(2, 3, 6));
let callCount = 0;
const memoizedFn = memoize(function (a, b) {
  callCount += 1;
  return a + b;
});
console.log(memoizedFn(0, 0));
console.log(memoizedFn(0, 0));
console.log(callCount);
addTwoPromises(Promise.resolve(2), Promise.resolve(2)).then(console.log);

const result = [];
const fn = (x) => x * 5;
const args = [2],
  t = 20,
  cancelT = 50;
const start = performance.now();
const log = (...argsArr) => {
  const diff = Math.floor(performance.now() - start);
  result.push({ time: diff, returned: fn(...argsArr) });
};

const cancel = cancellable(log, args, t);
const maxT = Math.max(t, cancelT);

setTimeout(() => {
  cancel();
}, cancelT);

setTimeout(() => {
  console.log(result); // [{"time":20,"returned":10}]
}, maxT + 1);

const limited = timeLimit((t) => new Promise((res) => setTimeout(res, t)), 100);
limited(150).catch(console.log); // "Time Limit Exceeded" at t=100ms

var obj = new TimeLimitedCache();
console.log("set1", obj.set(1, 2, 200));
console.log("set2", obj.set(10, 20, 400));
setTimeout(() => console.log("count1", obj.count()), 50);
setTimeout(() => console.log("count1", obj.count()), 100);
setTimeout(() => console.log("count1", obj.count()), 300);
setTimeout(() => console.log("count1", obj.count()), 500);

let start2 = Date.now();
function log2(...inputs) {
  console.log([Date.now() - start2, inputs]);
}
const dlog = debounce(log2, 50);
setTimeout(() => dlog(1), 50);
setTimeout(() => dlog(2), 75);

const promise = promiseAll([
  () => new Promise((resolve) => setTimeout(() => resolve(4), 50)),
  () => new Promise((resolve) => setTimeout(() => resolve(10), 150)),
  () => new Promise((resolve) => setTimeout(() => resolve(16), 100)),
]);
promise.then(console.log); // [42]
console.log("Object isEmpty:", isEmpty({ x: 5, y: 42 }));

var arr = [1, 9, 6, 3, 2],
  size = 3;
console.log(chunk(arr, size));

var sortArr = [{ x: 1 }, { x: 0 }, { x: -1 }],
  sortFn = (d) => d.x;
console.log(sortBy(sortArr, sortFn));

console.log(
  join(
    [
      { id: 1, r: 67, h: 83, d: 2 },
      { id: 2, f: 84, o: 46, l: 7 },
    ],
    [{ id: 1, c: 70, w: 1 }]
  )
);
