<template>
  <div id="map-container" ref="mapContainer" style="width: 100%; height: 400px;"></div>
</template>

<script>
import { onMounted, onBeforeUnmount, ref } from 'vue';
import trainStationData from '@/assets/train_station_data.json'; // 导入 JSON 文件

/* global AMap */

export default {
  name: 'MapComponent',
  setup() {
    const mapContainer = ref(null);
    // const logisticsInfoList = ref([
    //   { latitude: '23.129152403638752', longitude: '113.42775362698366' },
    //   { latitude: '30.454012', longitude: '114.42659' },
    //   { latitude: '31.93182', longitude: '118.633415' },
    //   { latitude: '31.035032', longitude: '121.611504' }
    // ]);

    let map = null;
    let intervalId = null;
    let marker = null;
    let isMapInitialized = false;
    const driving = ref(null);

    
    // 初始化地图
    const initMap = (longitude, latitude) => {
      map = new AMap.Map(mapContainer.value, {
        zoom: 15,
        center: [longitude, latitude],
      });
      AMap.plugin(['AMap.ToolBar', 'AMap.Geolocation', 'AMap.Driving'], function() {
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
            addTrainStationMarkers(); // 添加车站标记
          } else {
            console.error('定位失败:', result);
          }
        });

        const drivingOptions = {
          policy: AMap.DrivingPolicy.LEAST_TIME,
        };
        
        // 初始化驾车服务
        driving.value = new AMap.Driving({
          map: map.value,
          ...drivingOptions,
        });

        // 添加驾车服务回调
        driving.value.on('complete', drivingCallback);

        // 规划路线
        searchRoute();
      });

      // 在地图初始化后添加车站标记
      addTrainStationMarkers();
      isMapInitialized = true;
    };

        // 规划路线
    const searchRoute = () => {
      if (driving.value) {
        const startLngLat = new AMap.LngLat(116.379018, 39.865026); // 起点
        const endLngLat = new AMap.LngLat(116.42732, 39.903752);   // 终点
        driving.value.search(startLngLat, endLngLat);
      }
    };

    // 路线规划完成后的回调函数
    const drivingCallback = (result) => {
      console.log('路线规划完成:', result);
    };

    // 添加车站标记
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

// 定期更新当前位置并发送到服务器
const updateLocation = (driverId) => {
    if (navigator.geolocation) {
        intervalId = setInterval(() => {
            navigator.geolocation.getCurrentPosition(
                (position) => {
                    const longitude = position.coords.longitude;
                    const latitude = position.coords.latitude;

                    // 第一次获取位置时初始化地图，否则仅更新标记位置
                    if (!isMapInitialized) {
                        initMap(longitude, latitude);
                    } else if (marker) {
                        marker.setPosition([longitude, latitude]); // 更新标记位置
                    }

                    // 发送 GPS 数据到后端
                    fetch("http://localhost:8080/updateLocation", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify({
                            id: driverId,        // 驾驶员ID
                            role: "driver",      // 角色字段: driver
                            latitude: latitude,
                            longitude: longitude,
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
        }, 1000); // 每 1 秒更新一次位置
    } else {
        console.error("浏览器不支持地理定位");
    }
};


    onMounted(() => {
      updateLocation("driver1");
    });

    onBeforeUnmount(() => {
      if (intervalId) {
        clearInterval(intervalId);
      }
    });

    return {
      mapContainer,
    };
  },
};
</script>

<style scoped>
#map-container {
  width: 100%;
  height: 400px;
}
</style>
