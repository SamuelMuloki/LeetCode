class EventEmitter {
  constructor() {
    this.eventMap = new Map();
  }
  subscribe(event, cb) {
    var cbs = this.eventMap.get(event) || [];
    var jointArr = [...cbs, cb];
    this.eventMap.set(event, jointArr);

    return {
      unsubscribe: () => {
        var maps = this.eventMap.get(event);

        if (maps.length > 1) {
          maps.shift();
          this.eventMap.set(event, maps);
        } else {
          this.eventMap.delete(event);
        }
      },
    };
  }

  emit(event, args = []) {
    return this.eventMap.has(event)
      ? this.eventMap.get(event).map((fn) => fn.apply(null, args))
      : [];
  }
}

/**
 * const emitter = new EventEmitter();
 *
 * // Subscribe to the onClick event with onClickCallback
 * function onClickCallback() { return 99 }
 * const sub = emitter.subscribe('onClick', onClickCallback);
 *
 * emitter.emit('onClick'); // [99]
 * sub.unsubscribe(); // undefined
 * emitter.emit('onClick'); // []
 */
module.exports = { EventEmitter };
