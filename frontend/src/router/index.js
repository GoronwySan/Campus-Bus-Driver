import { createRouter, createWebHistory } from 'vue-router'; // Vue Router 4 的导入方式
import LoginPage from '../views/login.vue'; // 确保路径正确
import HomePage from '../views/HomePage.vue';
import Driver_Page from '@/views/driver/driver_Page.vue';
import Driver_Info from '@/views/driver/driver_Info.vue';


const routes = [
    {
        path: '/',
        redirect: '/login', // 访问根路径时重定向到 /login
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginPage,
    },
    {
        path: '/home',
        name: 'Home',
        component: HomePage,
    },
    {
        path: '/driver',
        name: 'Driver',
        component:Driver_Page
    },
    {
        path: '/driverInfo',
        name: 'DriberInfo',
        component:Driver_Info,
    }
    
];


const router = createRouter({
    history: createWebHistory(), // 使用 history 模式
    routes,
});

export default router;
