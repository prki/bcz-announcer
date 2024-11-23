import { reactive } from 'vue';

// [TODO] Ideally, these should store the cards in general.
// For our purposes, storing the count is fine enough so that create component
// gets info on whether it can create even.
export const calloutStore = reactive({
  cardCount: 0,
  maxCards: 12,

  increment() {
    if (this.cardCount >= this.maxCards) {
      console.log("[ERROR] Attempted to increment calloutStore cardCount, but was already at max!");
    }

    this.cardCount += 1;
  },

  decrement() {
    if (this.cardCount <= 0) {
      console.log("[ERROR] Attempted to decrement calloutStore cardCount, but was already 0");
      return;
    }

    this.cardCount -= 1;
  }
});
