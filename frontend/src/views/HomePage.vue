<script setup>
import router from '@/router';

function todriver() {
    router.push('/driver-0');
}

</script>


<template>
    <div>

        <h1>欢迎来到首页！</h1>
        <button @click="todriver" class="todriverButton">司机</button>

    </div>
</template>


<style scoped>
.todriverButton {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 100px;
    height: 52px;
    border-radius: 20px 20px 20px 20px;

    text-align: center;
    background: linear-gradient(to right,
            rgb(113, 65, 168),
            rgba(44, 114, 241, 1));
    color: floralwhite;
    line-height: 52px;
    cursor: pointer;
    font-size: 14px;
    user-select: none;
}
</style>