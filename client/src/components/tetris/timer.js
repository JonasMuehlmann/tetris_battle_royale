/***
 * SELF-ADJUSTING TIMER
 * DELAY IS ADAPTED TO THE ACTUALLY ELAPSED TIME
 */
export class Timer {
  constructor(interval = 1000) {
    this.interval = interval
  }

  run(callback) {
    this.expected = Date.now() + this.interval

    const step = () => {
      callback?.()

      const dt = Date.now() - this.expected
      this.expected += this.interval
      this.clear = setTimeout(step, Math.max(0, this.interval - dt))
    }

    this.clear = setTimeout(step, this.interval)
  }

  stop() {
    clearTimeout(this.clear)
  }
}