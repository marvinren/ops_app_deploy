<template>
  <div class="app-mgr-container">
    <el-card shadow="hover">
      <div class="system-menu-search mb15">
        <el-input size="default" placeholder="请输入应用名称" style="max-width: 180px"> </el-input>
        <el-button size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search />
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="success" class="ml10" @click="openDialog">
          <el-icon>
            <ele-FolderAdd />
          </el-icon>
          新增应用
        </el-button>
      </div>
      <el-table :data="applications">
        <el-table-column prop="name" label="应用名称" width="180"></el-table-column>
        <el-table-column prop="components" label="应用组件数量" width="180"></el-table-column>
        <el-table-column prop="admin" label="管理员" width="180"></el-table-column>
        <el-table-column prop="createAt" label="创建时间" width="180"></el-table-column>
        <el-table-column prop="createBy" label="创建人" width="180"></el-table-column>
        <el-table-column prop="description" label="应用描述"></el-table-column>
        <el-table-column fixed="right" label="操作" width="180">
          <template #default>
            <el-button type="text" size="small">新建组件</el-button>
            <el-button type="text" size="small" @click="app_info">组件详情</el-button>
            <el-button type="text" size="small" @click="app_conf">配置详情</el-button>
            <el-button type="text" size="small">编辑</el-button>
            <el-button type="text" size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog
        v-model="editDialogVisible"
        title="应用"
        width="30%"
        :before-close="handleClose"
    >
      <el-form :model="app" label-width="120px">
        <el-form-item label="应用名称">
          <el-input v-model="app.name" />
        </el-form-item>
        <el-form-item label="应用描述">
          <el-input v-model="app.description" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleClose">
          保存
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const editDialogVisible = ref(false)
const applications = reactive([
  {"id": 1, "name": "aiops-application-system", "description": "aiops微服务应用", "admin": "renzq", "components": 5, "createAt": "2022-01-02 00:00:00", "createBy": "renzq"},
  {"id": 2, "name": "devops-application-system", "description": "devops微服务应用", "admin": "renzq", "components": 8, "createAt": "2022-01-02 00:00:00", "createBy": "renzq"},

])
const app = reactive({
  name: "",
  description: "",
})

const openDialog = () => {
  editDialogVisible.value = true
}

const handleClose = () => {
  editDialogVisible.value = false
}

const app_info = () =>{
  router.push({path: "/app/appinfo"})
}

const app_conf = () => {
  router.push({path: "/app/appconf"})
}
</script>