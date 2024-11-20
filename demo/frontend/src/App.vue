<template>
  <div id="app">
    <div class="top-bar">
      <span>在线 <strong>1</strong> 人</span>
      <span>中山大学 (1号车)</span>
    </div>
    <div class="main-container">
      <!-- 地图部分 -->
      <div class="map-section">
        <MapComponent />
      </div>

      <!-- 信息部分 -->
      <div class="info-section">
        <div class="stats">
          <!-- 付款人数 -->
          <PaymentCount />
        </div>
      </div>
      <!-- 呼叫人数界面 -->
      <div class="info-section">
        <ErrorBoundary>
          <CallCount />
        </ErrorBoundary>
      </div>
    </div>
    <RouteEditor
      :polylineEditor="polylineEditor"
      :polylines="polylines"
      @remove-polyline="removePolyline"
    />
    <!-- 底部栏 -->
    <div class="footer-bar">
      <div class="footer-buttons">
        <ErrorBoundary>
          <SelfScanButton />
        </ErrorBoundary>
        <ErrorBoundary>
          <SafetyAlertButton />
        </ErrorBoundary>
      </div>
    </div>
  </div>
</template>

<script>
import { createPinia } from "pinia";
import MapComponent from './components/MapComponent.vue';
import SafetyAlertButton from '@/components/SafetyAlertButton.vue';
import SelfScanButton from '@/components/SelfScanButton.vue';
import PaymentCount from '@/components/PaymentCount.vue';
import CallCount from '@/components/CallCount.vue';
import ErrorBoundary from '@/components/ErrorBoundary.vue'; // 通用错误边界组件



export default {
  name: 'App',
  components: {
    MapComponent,
    SafetyAlertButton,
    SelfScanButton,
    PaymentCount,
    CallCount,
    ErrorBoundary,
  },
  created() {
    // 初始化 Pinia（确保在项目中已注册 Pinia）
    this.$pinia = createPinia();
  },
};
</script>

<style scoped>
/* #app {
  display: flex;
  flex-direction: column;
  height: 100vh;
} */
/* 顶部栏样式 */
.top-bar {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #2c3e50;
  color: #fff;
  font-size: 14px;
}

/* 主容器样式 */
.main-container {
  display: flex;
  flex-direction: row;
  height: 75vh;
}

/* 地图部分样式 */
.map-section {
  flex: 3;
  border-right: 2px solid #ddd;
}

/* 信息部分样式 */
.info-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: #ffffff;
  border: 1px solid #dcdcdc;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* 标题样式 */
.info-section h3 {
  margin-bottom: 10px;
  font-size: 18px;
  font-weight: bold;
  color: #333333;
}

/* 文本样式 */
.info-section p {
  margin: 5px 0;
  font-size: 16px;
  color: #666666;
}

/* 按钮样式 */
.action-btn {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  margin-top: 10px;
  transition: background-color 0.3s ease;
}

.action-btn:hover {
  background-color: #0056b3;
}

/* 数字统计样式 */
.stats {
  text-align: center;
  margin-bottom: 20px;
}

.current-count {
  font-size: 60px;
  color: #32cd32;
  font-weight: bold;
}

/* 按钮区域样式 */
.buttons {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

/* 底部栏样式 */
.footer-bar {
  display: flex;
  justify-content: flex-end; /* 按钮从右到左排列 */
  align-items: center;
  background-color: #f8f8f8;
  padding: 10px 20px;
  border-top: 1px solid #ddd;
}

/* 按钮样式 */
.footer-buttons > * {
  margin-left: 10px; /* 按钮之间的间距 */
}

.footer-buttons button {
  background-color: #007bff;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.footer-buttons button:hover {
  background-color: #0056b3;
}
</style>
