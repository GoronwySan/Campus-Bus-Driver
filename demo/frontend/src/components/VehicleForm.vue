<template>
    <div>
      <!-- 按钮 -->
      <button @click="openForm" class="form-button">处理工作班次</button>
  
      <!-- 信息填写窗口 -->
      <div v-if="showForm" class="form-modal">
        <div class="form-content">
          <h3>填写工作班次信息</h3>
          <form @submit.prevent="submitForm">
            <div class="form-item">
              <label for="route">路线路径:</label>
              <input v-model="formData.route" id="route" type="text" placeholder="请输入路线路径" required />
            </div>
            <div class="form-item">
              <label for="plate">车牌号:</label>
              <input v-model="formData.plate" id="plate" type="text" placeholder="请输入车牌号" required />
            </div>
            <div class="form-item">
              <label for="shiftAction">车辆状态:</label>
              <select v-model="formData.shiftAction" id="shiftAction" required>
                <option value="" disabled>请选择状态</option>
                <option value="start">正常运营</option>
                <option value="end">试通行</option>
              </select>
            </div>
            <!-- <div v-if="formData.shiftAction === 'end'" class="form-item">
              <label for="numPeople">下班人数:</label>
              <input v-model="formData.numPeople" id="numPeople" type="number" placeholder="请输入下班人数" required />
            </div> -->
            <div class="form-actions">
              <button type="submit" class="submit-button">提交</button>
              <button type="button" @click="closeForm" class="cancel-button">取消</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        showForm: false, // 控制弹窗显示与否
        formData: {
          driverID: "",    // 工号
          route: "",       // 路线路径
          plate: "",       // 车牌号
        },
      };
    },
    methods: {
      // 打开表单
      openForm() {
        this.showForm = true;
      },
      // 关闭表单
      closeForm() {
        this.showForm = false;
        this.resetForm();
      },
      // 重置表单
      resetForm() {
        this.formData = {
          driverID: "",
          route: "",
          plate: "",
          vehicle_status: "",
        };
      },
      // 提交表单
      async submitForm() {
        try {
          let endpoint = "";
          let method = "";
          let requestBody = {};
  
          // 根据操作选择不同的请求路径和请求体
          if (this.formData.shiftAction === "start") {
            endpoint = "http://localhost:3456/start";
            method = "POST";
            requestBody = {
              driver_id: this.formData.driverID,
              vehicle_no: this.formData.plate,
              route: this.formData.route,
            };
          } else if (this.formData.shiftAction === "end") {
            endpoint = "http://localhost:3456/end";
            method = "POST";
            requestBody = {
              driver_id: this.formData.driverID,
              vehicle_no: this.formData.plate,
              num_people: this.formData.numPeople,
            };
          }
  
          const response = await fetch(endpoint, {
            method: method,
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(requestBody),
          });
  
          const result = await response.json();
  
          if (response.ok) {
            // 信息提交成功
            alert("操作成功！");
            window.location.href = "/shift-detail"; // 替换为目标页面路径
          } else {
            // 错误处理
            alert(result.message || "操作失败，请检查输入！");
          }
        } catch (error) {
          console.error("提交失败:", error);
          alert("提交失败，请稍后再试！");
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .form-button {
    padding: 10px 20px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .form-modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .form-content {
    background-color: white;
    padding: 20px;
    border-radius: 4px;
    width: 400px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  .form-item {
    margin-bottom: 15px;
  }
  .form-item label {
    display: block;
    margin-bottom: 5px;
  }
  .form-item input,
  .form-item select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  .form-actions {
    display: flex;
    justify-content: space-between;
  }
  .submit-button {
    padding: 10px 20px;
    background-color: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  .cancel-button {
    padding: 10px 20px;
    background-color: #dc3545;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  </style>
  