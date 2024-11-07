<template>
  <div id="container"></div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from 'vue';
import AMapLoader from '@amap/amap-jsapi-loader';
import trainStationData from '@/assets/train_station_data.json'; // 导入 JSON 文件

/* global AMap */

let map = null;
let marker = null;
let intervalId = null;
let isMapInitialized = false;
const driving = ref(null);

const addTrainStationMarkers = () => {
  if (!map) return;
  const labelsLayer = new AMap.LabelsLayer({ collision: true });

  trainStationData.forEach((station) => {
    const labelMarker = new AMap.LabelMarker({
      position: station.position,
      text: {
        content: station.name,
        style: {
          fontSize: 15,
          fillColor: '#fff',
          backgroundColor: 'blue',
          borderColor: '#ccc',
          borderWidth: 2,
          padding: [5, 10],
        },
      },
      icon: {
        image: require('@/assets/circle-icon.png'),
        size: [15, 15],
        anchor: 'center',
      },
    });
    labelsLayer.add(labelMarker);
  });

  map.add(labelsLayer);
};

const drivingCallback = (result) => {
  console.log('路线规划完成:', result);
};

const searchRoute = () => {
  if (driving.value) {
    const startLngLat = new AMap.LngLat(116.379018, 39.865026);
    const endLngLat = new AMap.LngLat(116.42732, 39.903752);
    driving.value.search(startLngLat, endLngLat);
  }
};

const initMap = (longitude, latitude) => {
  map = new AMap.Map("container", {
    zoom: 15,
    center: [longitude, latitude],
  });

  AMap.plugin(['AMap.ToolBar', 'AMap.Geolocation', 'AMap.Driving'], () => {
    const toolbar = new AMap.ToolBar();
    map.addControl(toolbar);

    const geolocation = new AMap.Geolocation({
      enableHighAccuracy: true,
      timeout: 10000,
      buttonOffset: new AMap.Pixel(10, 20),
      zoomToAccuracy: true,
    });
    map.addControl(geolocation);

    geolocation.getCurrentPosition((status, result) => {
      if (status === 'complete') {
        console.log('定位成功:', result);
        addTrainStationMarkers();
      } else {
        console.error('定位失败:', result);
      }
    });

    driving.value = new AMap.Driving({
      map,
      policy: AMap.DrivingPolicy.LEAST_TIME,
    });

    driving.value.on('complete', drivingCallback);
    searchRoute();
  });

  addTrainStationMarkers();
  isMapInitialized = true;
};

const updateLocation = (driverId) => {
  if (navigator.geolocation) {
    intervalId = setInterval(() => {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { longitude, latitude } = position.coords;

          if (!isMapInitialized) {
            initMap(longitude, latitude);
          } else if (marker) {
            marker.setPosition([longitude, latitude]);
          }

          fetch("http://localhost:8080/updateLocation", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              id: driverId,
              role: "driver",
              latitude,
              longitude,
              timestamp: new Date().toISOString(),
            }),
          })
            .then((response) => response.text())
            .then((data) => console.log("服务器响应:", data))
            .catch((error) => console.error("请求错误:", error));
        },
        (error) => {
          console.error("无法获取位置", error);
        },
        { enableHighAccuracy: true, maximumAge: 0, timeout: 10000 }
      );
    }, 1000);
  } else {
    console.error("浏览器不支持地理定位");
  }
};

onMounted(() => {
  window._AMapSecurityConfig = {
    securityJsCode: "", //安全密钥
  };
  AMapLoader.load({
    key: "", // Web端开发者Key
    version: "2.0",
    plugins: ["AMap.Scale"],
  })
    .then(() => {
      updateLocation("driver1");
    })
    .catch((e) => {
      console.log(e);
    });
});

onUnmounted(() => {
  map?.destroy();
  if (intervalId) {
    clearInterval(intervalId);
  }
});
</script>

<style scoped>
#container {
  width: 100%;
  height: 800px;
}
</style>
