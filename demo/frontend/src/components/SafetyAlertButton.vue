<template>
    <button @click="playSafetyAlertAudio" class="safety-alert-btn">安全提醒</button>
  </template>
  
  <script>
  import { useAudioStore } from "@/stores/audioStore";
  
  export default {
    name: "SafetyAlertButton",
    data() {
      return {
        audioSrc: require("@/assets/safety-alert.mp3"), // 安全提醒音频文件路径
      };
    },
    methods: {
      playSafetyAlertAudio() {
        try {
          const audioStore = useAudioStore(); // 获取 Pinia Store 实例
          audioStore.playAudio(this.audioSrc); // 调用 Store 中的播放方法
          console.log("播放安全提醒音频");
        } catch (error) {
          console.error("播放安全提醒音频失败:", error);
          alert("安全提醒音频无法播放，请检查资源文件！");
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .safety-alert-btn {
    background-color: #28a745;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  .safety-alert-btn:hover {
    background-color: #1e7e34;
  }
  </style>
  