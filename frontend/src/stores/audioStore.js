import { defineStore } from "pinia";

export const useAudioStore = defineStore("audio", {
    state: () => ({
        currentAudio: null, // 当前播放的音频实例
    }),
    actions: {
        playAudio(audioSrc) {
            // 如果有音频在播放，停止并重置
            if (this.currentAudio && !this.currentAudio.paused) {
                this.currentAudio.pause();
                this.currentAudio.currentTime = 0;
            }

            // 播放新的音频
            this.currentAudio = new Audio(audioSrc);
            this.currentAudio.play();
        },
    },
});
