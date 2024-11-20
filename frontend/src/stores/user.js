import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        userAccount: null, // 用戶帳號
    }),
    actions: {
        setUserAccount(account) {
            this.userAccount = account; // 更新用戶帳號
        },
        clearUserAccount() {
            this.userAccount = null; // 清空用戶帳號
        },
    },
});
