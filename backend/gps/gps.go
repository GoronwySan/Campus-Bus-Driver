package gps

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// Driver 代表驾驶员的基本信息
type Driver struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GPSModule 用于管理驾驶员的信息和位置信息
type GPSModule struct {
	drivers map[string]*Driver
	mu      sync.Mutex
}

// NewGPSModule 创建一个 GPSModule 实例
func NewGPSModule() *GPSModule {
	return &GPSModule{
		drivers: make(map[string]*Driver),
	}
}

// CreateDriver 创建一个新的驾驶员对象
func (g *GPSModule) CreateDriver(id string, latitude, longitude float64) (*Driver, error) {
	if id == "" {
		return nil, errors.New("driver ID cannot be empty")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	driver := &Driver{
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
	}
	g.drivers[id] = driver
	return driver, nil
}

// GetAllDrivers 获取所有驾驶员的位置信息
func (g *GPSModule) GetAllDrivers() []*Driver {
	g.mu.Lock()
	defer g.mu.Unlock()

	drivers := make([]*Driver, 0, len(g.drivers))
	for _, driver := range g.drivers {
		drivers = append(drivers, driver)
	}
	return drivers
}

// UpdateDriverLocation 更新驾驶员的位置信息
func (g *GPSModule) UpdateDriverLocation(id string, latitude, longitude float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	driver, exists := g.drivers[id]
	if !exists {
		return errors.New("driver not found")
	}

	driver.Latitude = latitude
	driver.Longitude = longitude
	return nil
}

// Handler 接收来自前端的请求并处理 GPS 信息
func (g *GPSModule) Handler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID        string  `json:"id"`
		Role      string  `json:"role"` // "driver" or "passenger"
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}

	if requestData.Role == "driver" {
		// 更新驾驶员的位置信息
		err = g.UpdateDriverLocation(requestData.ID, requestData.Latitude, requestData.Longitude)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Driver location updated successfully")
	} else if requestData.Role == "passenger" {
		// 处理乘客信息（可扩展逻辑）
		fmt.Fprintf(w, "Passenger location received")
	} else {
		http.Error(w, "invalid role", http.StatusBadRequest)
	}

	// fmt.Printf("Received GPS data: ID %s, Role %s, Latitude %.6f, Longitude %.6f",
	// 	requestData.ID, requestData.Role, requestData.Latitude, requestData.Longitude)
}

// GetAllDriversHandler 公共接口，返回所有驾驶员的位置信息
func (g *GPSModule) GetAllDriversHandler(w http.ResponseWriter, r *http.Request) {
	drivers := g.GetAllDrivers()
	json.NewEncoder(w).Encode(drivers)
}

// DeleteDriver 删除一个驾驶员对象
func (gps *GPSModule) DeleteDriver(id string) error {
	if _, exists := gps.drivers[id]; !exists {
		return fmt.Errorf("驾驶员 %s 不存在", id)
	}
	delete(gps.drivers, id)
	fmt.Printf("驾驶员 %s 已删除\n", id)
	return nil
}
