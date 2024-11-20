<template>
    <button @click="playSelfScanAudio" class="self-scan-btn">自觉扫码</button>
  </template>
  
  <script>
  import { useAudioStore } from "@/stores/audioStore";
  
  export default {
    name: "SelfScanButton",
    data() {
      return {
        audioSrc: require("@/assets/self-scan.mp3"), // 自觉扫码音频文件路径
      };
    },
    methods: {
      playSelfScanAudio() {
        try {
          const audioStore = useAudioStore(); // 获取 Pinia Store 实例
          audioStore.playAudio(this.audioSrc); // 调用 Store 中的播放方法
          console.log("播放自觉扫码音频");
        } catch (error) {
          console.error("播放自觉扫码音频失败:", error);
          alert("自觉扫码音频无法播放，请检查资源文件！");
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .self-scan-btn {
    background-color: #007bff;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  .self-scan-btn:hover {
    background-color: #0056b3;
  }
  </style>
  