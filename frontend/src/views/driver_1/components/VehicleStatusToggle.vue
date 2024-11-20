<template>
    <div class="vehicle-status-toggle">
      <label for="vehicle-status">状态：</label>
      <select id="vehicle-status" v-model="selectedStatus" @change="updateStatus">
        <option value="normal">正常运营</option>
        <option value="trial">试通行</option>
      </select>
      <span class="status-display">当前状态：{{ statusText }}</span>
    </div>
  </template>
  
  <script>
  export default {
    name: "VehicleStatusToggle",
    data() {
      return {
        selectedStatus: "normal", // 默认状态
      };
    },
    computed: {
      statusText() {
        return this.selectedStatus === "normal" ? "正常运营" : "试通行";
      },
    },
    methods: {
      updateStatus() {
        // 触发地图上的状态更新
        this.$emit("status-change", this.selectedStatus);
      },
    },
  };
  </script>
  
  <style scoped>
  .vehicle-status-toggle {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  
  .vehicle-status-toggle select {
    padding: 5px;
    border-radius: 5px;
    border: 1px solid #ccc;
  }
  
  .status-display {
    font-weight: bold;
    color: #007bff;
  }
  </style>
  