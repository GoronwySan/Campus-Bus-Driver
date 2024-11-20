<template>
  <div>
    <button @click="startNewRoute" v-if="!editingNewRoute">新增路线</button>
    <button @click="endEditing(true)" v-if="editingNewRoute">保存并结束新增</button>
    <button @click="endEditing(false)" v-if="editingNewRoute">取消新增</button>

    <button @click="startEditingRoutes" v-if="!editingAllRoutes">编辑现有路线</button>
    <button @click="endEditingRoutes" v-if="editingAllRoutes">结束编辑</button>
    <button @click="saveEditedPolyline" v-if="editingAllRoutes">保存编辑</button>

    <div v-if="editingAllRoutes">
      <div v-for="(polyline, index) in polylines" :key="index">
        <button @click="startEditingPolyline(polyline)">编辑路线 {{ index + 1 }}</button>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  props: {
    polylineEditor: Object,
    polylines: Array,
  },
  data() {
    return {
      editingNewRoute: false,
      editingAllRoutes: false,
      newPolyline: null,
    };
  },
  methods: {
    startNewRoute() {
      if (!this.polylineEditor) return;

      const currentPolyline = this.polylineEditor.getTarget();
      if (currentPolyline) {
        const userConfirmed = window.confirm("当前正在编辑的路线尚未保存，是否新建路线？");
        if (!userConfirmed) {
          console.log("用户取消切换路线");
          return;
        }
        this.editingAllRoutes = false;
        this.polylineEditor.setTarget(null);
      } else {
        this.polylineEditor.setTarget();
        this.polylineEditor.open();
      }
      this.editingNewRoute = true;
    },
    endEditing(save) {
      try {
        this.newPolyline = this.polylineEditor?.getTarget();

        if (!this.newPolyline) {
          console.warn("没有找到新建的路线，可能已经被移除");
          this.editingNewRoute = false;
          return;
        }

        const path = this.newPolyline.getPath();
        this.polylineEditor.setTarget(null);
        this.polylineEditor.close();

        if (save) {
          const data = [{ path: path.map((point) => [point.lng, point.lat]) }];
          const blob = new Blob([JSON.stringify(data, null, 2)], { type: "application/json" });
          const url = URL.createObjectURL(blob);
          const link = document.createElement("a");
          link.href = url;
          link.download = "new_route.json";
          link.click();
          URL.revokeObjectURL(url);

          // 使用 $emit 通知父组件添加新路线
          this.$emit("add-polyline", this.newPolyline);
        } else {
          this.polylineEditor.setTarget(null);
          this.polylineEditor.close();
          this.$emit("remove-polyline", this.newPolyline); // 通知父组件删除路线
        }
      } catch (error) {
        console.error("在结束编辑时发生错误:", error);
      } finally {
        this.newPolyline = null;
        this.editingNewRoute = false;
      }
    },
    startEditingRoutes() {
      this.editingAllRoutes = true;
    },
    endEditingRoutes() {
      this.polylineEditor.setTarget(null);
      this.polylineEditor.close();
      this.editingAllRoutes = false;
    },
    saveEditedPolyline() {
      if (!this.polylineEditor) return;

      const targetPolyline = this.polylineEditor.getTarget();
      if (!targetPolyline) return;

      const updatedPath = targetPolyline.getPath();
      const userConfirmed = window.confirm("是否下载修改后的路线文件？");
      if (userConfirmed) {
        const data = [{ path: updatedPath.map((point) => [point.lng, point.lat]) }];
        const blob = new Blob([JSON.stringify(data, null, 2)], { type: "application/json" });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.href = url;
        link.download = "edited_route.json";
        link.click();
        URL.revokeObjectURL(url);
      }
      this.polylineEditor.setTarget(null);
      this.polylineEditor.close();
    },
    startEditingPolyline(polyline) {
      this.polylineEditor.setTarget(polyline);
      this.polylineEditor.open();
    },
  },
};
</script>
