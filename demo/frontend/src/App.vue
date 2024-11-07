<template>
  <div>
    <h1>我的高德地图</h1>
    <img alt="Vue logo" src="./assets/logo.png">
    <MapComponent />
  </div>
</template>

<script>
import MapComponent from './components/Map.vue';

export default {
  name: 'App',
  components: {
    MapComponent,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
