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

const limited = timeLimit((t) => new Promise(res => setTimeout(res, t)), 100);
limited(150).catch(console.log) // "Time Limit Exceeded" at t=100ms
