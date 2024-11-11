<template>
    <div class="payment-count-container">
      <h3 class="title">付款人数统计</h3>
      <div class="stats">
        <div class="stat-item">
          <span class="label">当天付款人数:</span>
          <span class="value">{{ currentCount }}</span>
        </div>
        <div class="stat-item">
          <span class="label">上次付款人数:</span>
          <span class="value">{{ lastCount }}</span>
        </div>
        <div class="stat-item">
          <span class="label">上上次付款人数:</span>
          <span class="value">{{ previousCount }}</span>
        </div>
      </div>
      <button class="clear-button" @click="clearCurrentCount">清空本次付款人数</button>
    </div>
  </template>
  
  <script>
  export default {
    name: "PaymentCount",
    data() {
      return {
        currentCount: 0, // 当前付款人数
        lastCount: 0, // 上次付款人数
        previousCount: 0, // 上上次付款人数
      };
    },
    methods: {
      clearCurrentCount() {
        try {
          this.previousCount = this.lastCount;
          this.lastCount = this.currentCount;
          this.currentCount = 0;
          console.log("付款人数已清空");
        } catch (error) {
          console.error("清空付款人数失败:", error);
          alert("清空付款人数时发生错误，请稍后再试！");
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .payment-count-container {
    width: 100%;
    max-width: 400px;
    margin: 20px auto;
    padding: 20px;
    border-radius: 8px;
    background-color: #f9f9f9;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    font-family: Arial, sans-serif;
  }
  
  .title {
    font-size: 1.5em;
    color: #333;
    text-align: center;
    margin-bottom: 20px;
    border-bottom: 2px solid #007bff;
    padding-bottom: 10px;
  }
  
  .stats {
    display: flex;
    flex-direction: column;
    gap: 10px;
    margin-bottom: 20px;
  }
  
  .stat-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .label {
    font-size: 1.2em;
    color: #555;
  }
  
  .value {
    font-size: 1.5em;
    font-weight: bold;
    color: #007bff;
  }
  
  .clear-button {
    display: block;
    width: 100%;
    padding: 10px 0;
    font-size: 1em;
    font-weight: bold;
    color: #fff;
    background-color: #ff4d4f;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .clear-button:hover {
    background-color: #d9363e;
  }
  </style>
  