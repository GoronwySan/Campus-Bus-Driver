// main.js
import { createApp } from "vue";
import App from "./App.vue";
import { createPinia } from "pinia";

const app = createApp(App);

// 创建 Pinia 实例
const pinia = createPinia();

// 注册 Pinia
app.use(pinia);

// 挂载 Vue 应用
app.mount("#app");
