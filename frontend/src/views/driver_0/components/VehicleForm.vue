<template>
  <div>
    <!-- 按钮 -->
    <button @click="openForm" class="form-button">上班</button>

    <!-- 信息填写窗口 -->
    <div v-if="showForm" class="form-modal">
      <div class="form-content">
        <h3>上班</h3>
        <form @submit.prevent="submitForm">
          <div class="form-item">
            <label for="route_id">路线编号:</label>
            <input v-model="formData.route_id" id="route_id" type="text" placeholder="请输入路线编号" required />
          </div>
          <div class="form-item">
            <label for="vehicleNo">车牌号:</label>
            <input v-model="formData.vehicleNo" id="vehicleNo" type="text" placeholder="请输入车牌号" required />
          </div>
          <div class="form-item">
            <label for="shiftAction">车辆状态:</label>
            <select v-model="formData.shiftAction" id="shiftAction" required>
              <option value="" disabled>请选择状态</option>
              <option value="start">正常运营</option>
              <option value="end">试通行</option>
            </select>
          </div>
          <div class="form-actions">
            <button type="submit" class="submit-button">提交</button>
            <button type="button" @click="closeForm" class="cancel-button">取消</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useUserStore } from "@/stores/user";

// 控制弹窗显示与否
const showForm = ref(false);

// 表单数据
const formData = ref({
  // driverID: "", // 工号
  route_id: "", // 路线路径
  vehicleNo: "", // 车牌号
  shiftAction: "", // 操作状态
  // shiftStart: "" // 上班時間
});

// 打开表单
const openForm = () => {
  showForm.value = true;
};

// 关闭表单并重置数据
const closeForm = () => {
  showForm.value = false;
  resetForm();
};

// 重置表单
const resetForm = () => {
  formData.value = {
    driverID: "",
    route_id: "",
    vehicleNo: "",
    shiftAction: "",
    shiftStart: ""
  };
};

//生成符合datetime格式的时间字符串
const getCurrentDatetime = () => {
  const now = new Date();
  const pad = (num) => num.toString().padStart(2, '0');
  return `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())} ${pad(now.getHours())}:${pad(now.getMinutes())}:${pad(now.getSeconds())}`;
};


// 提交表单
const submitForm = async () => {
  try {
    const userStore = useUserStore(); // 引入全局的 userStore
    const driverID = userStore.userAccount; // 獲取全局變量中的 driver_id

    let endpoint = "";
    let requestBody = {};

    endpoint = "http://localhost:8888/startWord";
    requestBody = {
      driver_id: driverID,
      vehicle_no: formData.value.vehicleNo,
      route_id: formData.value.route_id,
      shiftStart: getCurrentDatetime(),
    };

    const response = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestBody),
    });

    const result = await response.json();

    if (response.ok) {
      alert("操作成功！");
      closeForm();

      // 根据操作选择不同的请求路径和请求体
      if (formData.value.shiftAction === "") {
        console.error("null");
      }

      // 跳转界面
      // this.$router.push('/driver-1');
      window.location.href = "/shift-detail"; // 替换为目标页面路径

    } else {
      alert(result.message || "操作失败，请检查输入！");
    }
  } catch (error) {
    console.error("提交失败:", error);
    alert("提交失败，请稍后再试！");
  }
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
