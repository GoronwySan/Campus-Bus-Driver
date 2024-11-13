<template>
    <div>
      <h3>呼叫人数</h3>
      <p>今日呼叫人数: {{ callCount }}</p>
    </div>
  </template>
  
  <script>
  export default {
    name: "CallCount",
    data() {
      return {
        callCount: 0,
      };
    },
    methods: {
      incrementCallCount() {
        try {
          this.callCount++;
          console.log("呼叫人数增加，当前人数:", this.callCount);
        } catch (error) {
          console.error("增加呼叫人数失败:", error);
          alert("更新呼叫人数时发生错误，请稍后再试！");
        }
      },
    },
  };
  </script>
  