package main

import (
	"back/gps" // 引入 gps 包
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GPSData 结构体，用于接收前端的 GPS 数据
type GPSData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp string  `json:"timestamp"` // 或 time.Time, 取决于数据格式
}

// CORS 中间件
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://localhost:8081") // 设置允许的域
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")    // 允许的请求方法
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")          // 允许的请求头

		// 如果是预检请求（OPTIONS），则直接返回
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// handleGPSData 处理接收 GPS 数据的 POST 请求
func handleGPSData(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 响应头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 如果是预检请求，直接返回状态 200
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 仅支持 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "仅支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	var data GPSData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("请求数据解析错误: %v\n", err)
		http.Error(w, "请求数据格式错误", http.StatusBadRequest)
		return
	}

	fmt.Printf("接收到的 GPS 数据：纬度 %.6f，经度 %.6f，时间戳 %s\n", data.Latitude, data.Longitude, data.Timestamp)

	response := map[string]string{"message": "GPS 数据接收成功"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// 初始化 GPS 模块
	gpsModule := gps.NewGPSModule()

	// 示例：添加一个驾驶员对象
	gpsModule.CreateDriver("driver1", 34.0522, -118.2437)

	// 创建 ServeMux 路由
	mux := http.NewServeMux()

	// 设置路由，接收 GPS 数据的端点
	mux.HandleFunc("/api/gps", handleGPSData) // 接收 GPS 数据
	mux.HandleFunc("/createDriver", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing driver id", http.StatusBadRequest)
			return
		}
		latitude := 34.0522
		longitude := -118.2437
		driver, err := gpsModule.CreateDriver(id, latitude, longitude)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(driver)
	})

	// 用于获取所有驾驶员位置信息的接口
	mux.HandleFunc("/drivers", gpsModule.GetAllDriversHandler)

	// 用于接收并处理 GPS 信息的接口
	mux.HandleFunc("/updateLocation", gpsModule.Handler)

	// 使用 CORS 中间件
	corsHandler := enableCORS(mux)

	// 启动服务器
	fmt.Println("服务器正在运行于 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
