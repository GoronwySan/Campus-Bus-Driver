<script setup>
import { reactive, defineProps, defineEmits } from 'vue';

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  content: {
    type: String,
    default: ''
  }
});

const emit = defineEmits(['close']);

const closepopup = () => {
  emit('close');
};

// 用户信息
const user = reactive({
  avatar: require('@/assets/logo.png'), // 头像图片链接
  name: '张三',
  username: 'zhangsan',
  password: '******',
  gender: '男',
  phone: '1234567890',
  route: '1号线',
  licensePlate: '粤B12345'
});

// 用于备份用户数据
let userBackup = { ...user };

// 控制编辑模式
const isEditingMode = reactive({ value: false });

// 切换编辑模式
const toggleEditMode = () => {
  if (isEditingMode.value) {
    // 在编辑模式退出时备份数据
    userBackup = { ...user };
  }
  isEditingMode.value = !isEditingMode.value;
};

// 保存修改
const saveChanges = () => {
  console.log("保存修改:", user);
  isEditingMode.value = false; 
};

// 取消修改
const cancelChanges = () => {
  // 取消时恢复原来的数据
  Object.assign(user, userBackup);
  isEditingMode.value = false; 
};

</script>

<template>
  <div>
    <div v-if="props.visible" class="popup">
      <div class="popup-content">
        <div class="profile-container">
          <div class="avatar">
            <img :src="user.avatar" alt="头像" />
          </div>
          <div class="info">
            <p>
              <strong>姓名:</strong>
              <span v-if="!isEditingMode.value">{{ user.name }}</span>
              <input v-else v-model="user.name" type="text" />
            </p>
            <p>
              <strong>账号:</strong>
              <span v-if="!isEditingMode.value">{{ user.username }}</span>
              <input v-else v-model="user.username" type="text" />
            </p>
            <p>
              <strong>密码:</strong>
              <span v-if="!isEditingMode.value">{{ user.password }}</span>
              <input v-else v-model="user.password" type="password" />
            </p>
            <p>
              <strong>性别:</strong>
              <span v-if="!isEditingMode.value">{{ user.gender }}</span>
              <select v-else v-model="user.gender">
                <option value="男">男</option>
                <option value="女">女</option>
              </select>
            </p>
            <p>
              <strong>电话号:</strong>
              <span v-if="!isEditingMode.value">{{ user.phone }}</span>
              <input v-else v-model="user.phone" type="text" />
            </p>
            <p>
              <strong>路线:</strong>
              <span v-if="!isEditingMode.value">{{ user.route }}</span>
              <input v-else v-model="user.route" type="text" />
            </p>
            <p>
              <strong>车牌号:</strong>
              <span v-if="!isEditingMode.value">{{ user.licensePlate }}</span>
              <input v-else v-model="user.licensePlate" type="text" />
            </p>
          </div>

          <button @click="closepopup" class="closebutton">X</button>
          
          
          <button v-if="!isEditingMode.value" @click="toggleEditMode" class="modifybutton">
            修改信息
          </button>

          <button v-if="!isEditingMode.value" @click=" 1" class="logoutbutton">
            退出账号
          </button>

          
          <div v-if="isEditingMode.value" class="action-buttons">
            <button @click="saveChanges" class="savebutton">保存</button>
            
            <button @click="cancelChanges" class="cancelbutton">取消</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 80px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  position: relative;
}

.avatar img {
  user-select: none;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-bottom: 10px;
  cursor: pointer;
}

.info {
  text-align: left;
}

.info p {
  margin: 10px 0;
  font-size: 16px;
  color: #333;
}

.popup {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup-content {
  background: #fff;
  padding: 25px;
  border-radius: 30px;
  width: 600px;
  max-width: 90%;
  position: relative;
}

.closebutton {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 16px;
  color: #aaa;
  cursor: pointer;
}

.modifybutton {
  position: absolute;
  top: 87%;
  left: 30%;
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 20px 20px 20px 20px;
  cursor: pointer;
}

.modifybutton:hover {
  background-color: #45a049;
}

.logoutbutton{
  position: absolute;
  top: 87%;
  left: 55%;
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 20px 20px 20px 20px;
  cursor: pointer;
}

.logoutbutton:hover {
  background-color: #45a049; /* 鼠标悬停时的颜色变化 */
}

.action-buttons {
  margin-top: 20px;
  display: flex;
  justify-content: space-around;
}

.savebutton,
.cancelbutton {
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.savebutton {
  background-color: #4CAF50;
  color: white;
}

.cancelbutton {
  background-color: #f44336;
  color: white;
}

.savebutton:hover {
  background-color: #45a049;
}

.cancelbutton:hover {
  background-color: #e02a2a;
}
</style>
