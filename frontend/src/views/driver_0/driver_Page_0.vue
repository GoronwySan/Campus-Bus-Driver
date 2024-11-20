<script setup>
import MapComponent from '@/views/driver_0/components/MapComponent.vue';
</script>

<template>
  <div id="app">
    <div class="top-bar">
      <span>在线 <strong>1</strong> 人</span>
      <span>中山大学 (1号车)</span>
    </div>
    <div class="main-container">
      <MapComponent />
    </div>
  </div>
</template>

<style scoped>
.page-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  font-family: Arial, sans-serif;
  background-color: #f3f4f6;
  min-height: 100vh;
}

.page-title {
  color: #4a4a4a;
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}

.map-container {
  position: relative;
  height: 700px;
  width: 100%;
  max-width: 800px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  margin-bottom: 20px;
}

/* 顶部栏样式 */
.top-bar {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #2c3e50;
  color: #fff;
  font-size: 14px;
}

/* 地图部分样式 */
.map-section {
  flex: 3;
  border-right: 2px solid #ddd;
}



@media (min-width: 1024px) {
  .map-container {
    height: 600px;
  }
}
</style>
