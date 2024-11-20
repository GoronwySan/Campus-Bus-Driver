<template>
    <div class="page-container">
        <div id="container" class="map-container">
            <div class="info-container">
                <button class="btn" @click="toggleRoutes" style="margin-bottom: 5px">
                    {{ routesVisible ? '隐藏现有路线' : '显示现有路线' }}
                </button>
                <button class="btn" @click="toggleStations" style="margin-bottom: 5px">
                    {{ stationsVisible ? '隐藏站点' : '显示站点' }}
                </button>
            </div>
        </div>
    </div>
</template>


<script>
import AMapLoader from "@amap/amap-jsapi-loader";
import busStationData from "@/assets/bus_station_data.json";

import { useUserStore } from "@/stores/user";

/* global AMap */

export default {
    name: "MapComponent",
    data() {
        return {
            map: null, // 地图实例，用于存储高德地图对象
            marker: null, // 当前用户位置的标记
            intervalId: null, // 定时器 ID，用于位置更新
            isMapInitialized: false, // 地图是否已经初始化，用于避免重复初始化
            polylineEditor: null, // 折线编辑器实例，用于编辑路线
            polylines: [], // 路线的折线对象数组，存储所有绘制在地图上的路线
            polyline: null, // 当前正在操作的折线对象
            stationMarkers: [], // 存储所有站点的标记对象，用于显示或隐藏站点
            routesVisible: true, // 是否显示当前绘制的路线，用于控制路线的可见性
            stationsVisible: true, // 是否显示站点，用于控制站点标记的可见性
            editingNewRoute: false, // 是否正在新增路线，用于控制新增路线模式
            newPolyline: null, // 新建的折线对象，用于新增路线时的存储
            editingAllRoutes: false, // 是否正在编辑所有路线，用于控制编辑路线模式
            onlineCount: 1, // 假设初始在线人数
            drivers: [], // 存储从后端获取的驾驶员位置数据
            markers: [], // 存储地图上的标记
            dInfoVisible: false
        };
    },
    methods: {
        closeDInfo() {
            this.dInfoVisible = false;
        },
        showDriverInfo() {
            this.dInfoVisible = true;
        },
        initMap(longitude, latitude) {
            this.map = new AMap.Map("container", {
                zoom: 15,
                center: [longitude, latitude],
            });

            AMap.plugin(
                ["AMap.Geolocation", "AMap.Driving", "AMap.PolylineEditor"],
                () => {
                    const geolocation = new AMap.Geolocation({
                        enableHighAccuracy: true,
                        timeout: 10000,
                        buttonOffset: new AMap.Pixel(10, 20),
                        zoomToAccuracy: true,
                    });
                    this.map.addControl(geolocation);

                    geolocation.getCurrentPosition((status, result) => {
                        if (status === "complete") {
                            console.log("定位成功:", result);
                            this.addBusStationMarkers();
                        } else {
                            console.error("定位失败:", result);
                        }
                    });

                    this.polylineEditor = new AMap.PolylineEditor(this.map, this.polyline);
                }
            );

            this.addBusStationMarkers();
            this.loadAndDrawRoutes(); // 加载并绘制路线
            this.isMapInitialized = true;
        },
        /** 自动定位到校区 */
        autoLocateCampus() {
            // 校区的经纬度
            const campusCenter = [113.584845, 22.358088];
            if (!this.map) return;

            this.map.setZoomAndCenter(15, campusCenter); // 设置缩放级别和中心点
            console.log("已定位到校区中心:", campusCenter);
        },
        /** 显示/隐藏现有路线 */
        toggleRoutes() {
            this.routesVisible = !this.routesVisible;
            this.polylines.forEach((polyline) => {
                if (this.routesVisible) {
                    this.map.add(polyline);
                } else {
                    this.map.remove(polyline);
                }
            });
        },
        /** 显示/隐藏站点 */
        toggleStations() {
            this.stationsVisible = !this.stationsVisible;
            this.stationMarkers.forEach((marker) => {
                if (this.stationsVisible) {
                    this.map.add(marker);
                } else {
                    this.map.remove(marker);
                }
            });
        },
        /** 动态加载路线 */
        loadAndDrawRoutes() {
            const context = require.context('@/assets', false, /^\.\/route[0-9]+\.json$/); // 匹配以 route 开头的 JSON 文件
            const routes = [];
            console.log("匹配到的文件:", context.keys());

            context.keys().forEach((fileName) => {
                const fileData = context(fileName);
                console.log("加载的文件内容:", fileName, fileData);

                // 确保文件内容格式正确
                if (Array.isArray(fileData)) {
                    fileData.forEach((route) => {
                        if (route && route.path) {
                            routes.push(route.path);
                        }
                    });
                }
            });

            // 绘制所有路径到地图
            routes.forEach((path, index) => {
                const polyline = new AMap.Polyline({
                    path: path, // 路线点数组
                    strokeColor: this.getRouteColor(index), // 根据索引设置颜色
                    strokeWeight: 6,
                });

                this.map.add(polyline); // 添加到地图
                this.polylines.push(polyline); // 保存到 polylines 数组
                console.log(`绘制路径 ${index + 1} 成功`, path);
            });

            console.log("所有符合条件的路线已加载并绘制");
        },
        /** 添加站点标记 */
        addBusStationMarkers() {
            if (!this.map) return;
            busStationData.forEach((station) => {
                const labelMarker = new AMap.LabelMarker({
                    position: station.position,
                    text: {
                        content: station.name,
                        style: {
                            fontSize: 15,
                            fillColor: "#fff",
                            backgroundColor: "blue",
                            borderColor: "#ccc",
                            borderWidth: 2,
                            padding: [5, 10],
                        },
                    },
                    icon: {
                        image: require("@/assets/circle-icon.png"),
                        size: [15, 15],
                        anchor: "center",
                    },
                });
                this.stationMarkers.push(labelMarker);
                this.map.add(labelMarker);
            });
        },
        /** 路线颜色 */
        getRouteColor(index) {
            const colors = ['#FF0000', '#00FF00', '#0000FF', '#FFFF00', '#FF00FF']; // 可扩展颜色列表
            return colors[index % colors.length];
        },
        handleStatusChange(status) {
            console.log("状态已更新为：", status);
            // 更新地图上显示的状态
            this.updateMapStatus(status);
        },
        updateMapStatus(status) {
            // 示例：动态在地图上显示当前状态
            const statusDisplay = document.getElementById("map-status-display");
            if (!statusDisplay) {
                const newStatus = document.createElement("div");
                newStatus.id = "map-status-display";
                newStatus.style.position = "absolute";
                newStatus.style.top = "10px";
                newStatus.style.right = "10px";
                newStatus.style.background = "rgba(0, 0, 0, 0.5)";
                newStatus.style.color = "white";
                newStatus.style.padding = "5px 10px";
                newStatus.style.borderRadius = "5px";
                document.body.appendChild(newStatus);
                newStatus.innerText = `状态：${status === "normal" ? "正常运营" : "试通行"}`;
            } else {
                statusDisplay.innerText = `状态：${status === "normal" ? "正常运营" : "试通行"}`;
            }
        },
        // 管理员的功能
        addPolyline(newPolyline) {
            this.polylines.push(newPolyline); // 添加新路线
        },
        removePolyline(polyline) {
            const index = this.polylines.indexOf(polyline);
            if (index > -1) {
                this.polylines.splice(index, 1); // 移除指定的路线
            }
        },

        /** 更新位置 */
        updateLocation() {
            if (navigator.geolocation) {
                this.intervalId = setInterval(() => {
                    navigator.geolocation.getCurrentPosition(
                        (position) => {
                            const { longitude, latitude } = position.coords;

                            if (!this.isMapInitialized) {
                                this.initMap(longitude, latitude);
                            } else if (this.marker) {
                                this.marker.setPosition([longitude, latitude]);
                            }
                            // 调用发送位置信息到后端的方法
                            this.sendLocationToBackend(longitude, latitude);
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
        },
        // 发送位置信息到后端
        sendLocationToBackend(longitude, latitude) {
            const userStore = useUserStore(); // 引入全局的 userStore
            const driverID = userStore.userAccount; // 獲取全局變量中的 driver_id
            fetch("http://localhost:8888/updateLocation", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    id: driverID,
                    role: "driver", // 用户角色
                    latitude,
                    longitude,
                    timestamp: new Date().toISOString(), // 时间戳
                }),
            })
                .then((response) => response.text())
                .then((data) => console.log("服务器响应:", data))
                .catch((error) => console.error("请求错误:", error));
        },

        // 获取驾驶员数据
        async fetchDrivers() {
            try {
                const response = await fetch("http://localhost:8888/drivers");
                if (!response.ok) {
                    throw new Error("网络请求失败");
                }
                this.drivers = await response.json();

                // 验证驾驶员数据是否有效
                this.drivers = this.drivers.filter(
                    driver => typeof driver.latitude === "number" && typeof driver.longitude === "number"
                );

                this.updateMarkers(); // 更新地图上的标记
            } catch (error) {
                console.error("获取驾驶员位置失败:", error);
            }
        },

        // 在地图上显示驾驶员位置
        updateMarkers() {
            // 清除旧的标记
            this.markers.forEach(marker => this.map.remove(marker));
            this.markers = [];

            // 根据新的驾驶员数据添加标记
            this.drivers.forEach(driver => {
                const marker = new AMap.Marker({
                    position: [driver.longitude, driver.latitude], // 使用驾驶员的经纬度
                    map: this.map,
                    // icon: require('@/assets/driver-icon.png') // 引用自定义图标
                });

                this.markers.push(marker);
            });
        }
    },
    mounted() {
        window._AMapSecurityConfig = {
            securityJsCode: "bc6f966d4758af8f40837aa7560ada04", // 安全密钥
        };
        AMapLoader.load({
            key: "b97c0e27127e8ce02bbba2c585de79b1", // Web端开发者Key
            version: "2.0",
            plugins: ["AMap.Scale"],
        })
            .then(() => {
                this.updateLocation();
                this.fetchDrivers(); // 获取驾驶员数据
                // 可选：设置定时器定期刷新位置
                setInterval(() => {
                    this.fetchDrivers();
                }, 1000); // 每 1 秒刷新一次
            })
            .catch((e) => {
                console.log(e);
            });
    },
    beforeUnmount() {
        this.map?.destroy();
        if (this.intervalId) {
            clearInterval(this.intervalId);
        }
    },
};
</script>

<style scoped>
#container {
    width: 100%;
    height: 500px;
}

#map-wrapper {
    position: relative;
    width: 60%;
    height: 60%;
}

/* 输入卡片样式 */
.input-card {
    margin-top: 20px;
    padding: 10px;
    background-color: #f9f9f9;
    border: 1px solid #ddd;
    border-radius: 5px;
}

/* 地图顶部状态栏样式 */
.map-top-bar {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    background-color: rgba(255, 255, 255, 0.9);
    /* 半透明背景 */
    padding: 10px;
    z-index: 1000;
    /* 确保覆盖在地图之上 */
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.map-status-bar {
    position: absolute;
    top: 10px;
    left: 10px;
    display: flex;
    align-items: center;
    gap: 15px;
    background: rgba(255, 255, 255, 0.8);
    padding: 5px 15px;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
}

.online-count {
    font-size: 14px;
    color: #007bff;
}

/* 地图容器样式 */
#map-container {
    width: 100%;
    height: 60%;
}
</style>
<style scoped>
/* 地图容器样式 */
#map-wrapper {
    position: relative;
    width: 100%;
    height: 100px;
}

/* 地图顶部覆盖条 */
.map-top-bar {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    background-color: rgba(255, 255, 255, 0.9);
    /* 半透明背景 */
    padding: 10px;
    z-index: 1000;
    /* 确保覆盖在地图之上 */
    display: flex;
    align-items: center;
    justify-content: space-between;
}

/* 地图容器样式 */
#map-container {
    width: 100%;
    height: 100%;
    z-index: 1;
    /* 底层组件 */
}

/* 输入卡片样式 */
.input-card {
    margin-top: 20px;
    padding: 10px;
    background-color: #f9f9f9;
    border: 1px solid #ddd;
    border-radius: 5px;
}

/* 按钮样式 */
.btn {
    display: block;
    width: 100%;
    padding: 8px;
    background-color: #007bff;
    color: white;
    text-align: center;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.btn:hover {
    background-color: #0056b3;
}

.showDriverInfoButton {
    position: absolute;
    top: 20%;
    left: 85%;
    width: 52px;
    height: 52px;
    border-radius: 50%;

    text-align: center;
    background: rgba(113, 65, 168, 0.5);
    color: floralwhite;
    line-height: 52px;
    cursor: pointer;
    font-size: 14px;
    user-select: none;
    transform: translate(-50%, -50%);
}

.page-container {
    width: 95%;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    font-family: Arial, sans-serif;
    background-color: #f3f4f6;
    min-height: 100vh;
}

.page-title {
    color: #4a4a4a;
    font-size: 24px;
    margin-bottom: 20px;
    text-align: center;
}

.map-container {
    position: relative;
    height: 700px;
    width: 100%;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    margin-bottom: 20px;
}

.info-container {
    position: absolute;
    top: 10px;
    right: 10px;
    width: 260px;
    background-color: #ffffff;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
    color: #333;
    z-index: 10;
}

.info-container h4 {
    color: #008a6c;
    font-weight: 600;
}

.info-container p {
    color: #666;
    line-height: 1.6;
}

@media (min-width: 1024px) {
    .map-container {
        height: 600px;
    }
}
</style>