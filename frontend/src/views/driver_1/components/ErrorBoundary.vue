<template>
    <div>
      <slot v-if="!error" />
      <p v-else>组件加载时发生错误，请稍后再试。</p>
    </div>
  </template>
  
  <script>
  export default {
    name: "ErrorBoundary",
    data() {
      return {
        error: null,
      };
    },
    errorCaptured(err, vm, info) {
      console.error("错误边界捕获的错误:", err, info);
      this.error = err;
      return false; // 防止错误冒泡
    },
  };
  </script>
  