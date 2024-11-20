// src/services/loginService.js
import { useUserStore } from '@/stores/user';
import axios from "axios";

export default {
    data() {
        return {
            username: "", // 邮箱账号
            password: "", // 密码
        };
    },
    methods: {
        /**
         * 处理用户的登录请求
         *
         * 此函数会验证用户输入的账号和密码是否为空，并向服务器发送登录请求。
         * 若成功，则提示登录成功并进行后续处理；若失败，则给出错误信息提示。
         *
         * @returns {void} 无返回值
         * @throws {null} 如果请求失败或网络错误，在控制台和用户界面显示错误信息
         */
        handleLogin() {
            if (this.username === "" || this.password === "") {
                alert('账号或密码不能为空');
                return;
            }

            // 获取用户 Store
            const userStore = useUserStore();

            // 发送post请求
            axios.post("http://localhost:8888/api/login", {
                username: this.username,
                password: this.password,
            },
                {
                    timeout: 5000, // 设置超时时间，否则axios会一直等待
                })
                .then(response => {
                    if (response.data.code === 200) {
                        console.log("登陆成功：", response.data.data);

                        // 保存用户名到全局状态
                        userStore.setUserAccount(this.username);

                        alert('登录成功');
                        // 跳转界面
                        this.$router.push('/home');
                    } else {
                        alert('登录失败');
                        console.log("登陆失败：", response.data.data);
                    }
                })
                .catch(error => {
                    // 错误处理
                    if (error.response) {
                        // 请求已发送，服务器返回状态码，但状态码超出 2xx 范围
                        console.log("服务器错误:", error.response.data);
                        alert("登录失败，服务器错误：" + (error.response.data.message || "未知错误"));
                    } else if (error.request) {
                        // 请求已发送，但未收到响应
                        console.log("请求未收到响应:", error.request);
                        alert("登录失败，未收到服务器响应，请检查网络或服务器状态");
                    } else {
                        // 请求配置出错
                        console.log("请求配置错误:", error.message);
                        alert("登录失败，错误信息：" + error.message);
                    }
                });
        },
    },
};
