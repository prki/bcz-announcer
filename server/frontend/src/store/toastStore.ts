import { reactive } from 'vue';

export type Toast = {
  headerMessage: string,
  message: string,
  toastBg: string,  // CSS style for bg, e.g. "bg-danger"
  id: number
};

export const toastStore = reactive({
  toasts: [] as Toast[],

  // addToast should omit id in the type parameter
  addToast(toast: Omit<Toast, "id">) {
    const id = Date.now()
    this.toasts.push({...toast, id: id});

    const ttl = 5000;

    setTimeout(() => {
      // [TODO] using '!' to force TS to believe id exists. Add omit when creating.
      this.removeToast(id);
    }, ttl);
  },

  removeToast(id: number) {
    this.toasts = this.toasts.filter((toast) => toast.id !== id);
  }
});
