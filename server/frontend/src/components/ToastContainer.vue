<template>
  <div class="toast-container position-fixed top-0 end-0">
    <div v-for="(toast, idx) in toastStore.toasts" :key="idx" class="toast show" :class="toast.toastBg">
      <div class="toast-header">
        <strong class="me-auto">{{ toast.headerMessage }}</strong>
        <!-- [TODO] call close to delete from toaststore -->
        <button type="button" class="btn-close" @click="closeToast(toast)"></button>
      </div>
      <div class="toast-body">
        {{ toast.message }}
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { toastStore, Toast } from '../store/toastStore';

// [TODO] id! to force typescript to know it exists. However, this is not guaranteed with some changes.
// => use Omit<> on creation and force id to exist.
const closeToast = (toast: Toast) => {
  toastStore.removeToast(toast.id);
};
</script>

<style scoped>
</style>
