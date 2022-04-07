<template>
  <div class="app-appconf-container">
    <el-card shadow="hover">
      <div class="app-appconf-search mb15">
        <el-button size="default" type="default" class="ml10" @click="handleCancel">
          取消
        </el-button>
        <el-button size="default" type="default" class="ml10">
          部署
        </el-button>
        <el-button size="default" type="primary" class="ml10" >
          新增组件
        </el-button>
      </div>

      <el-row>
        <el-col :span="6">
          <div class="app-appconf-app-info">
            <el-descriptions title="应用信息" column="1" >
              <el-descriptions-item label="关联流水线">flow-2032144</el-descriptions-item>
              <el-descriptions-item label="创建人   ">admin</el-descriptions-item>
              <el-descriptions-item label="修改人  ">admin</el-descriptions-item>
              <el-descriptions-item label="创建时间">2022-03-05 13:00:12</el-descriptions-item>
              <el-descriptions-item label="修改时间">2022-03-05 13:00:12</el-descriptions-item>
            </el-descriptions>
          </div>
        </el-col>
        <el-col :span="18">
          <el-table :data="app_conf_datum" style="width: 100%">
            <el-table-column prop="name" label="组件名称" width="180"/>
            <el-table-column prop="category" label="组件类型" width="180"/>
            <el-table-column prop="order" label="部署顺序" width="180"/>
            <el-table-column prop="description" label="描述"/>
            <el-table-column fixed="right" label="操作" width="120">
              <template #default>
                <el-button type="text" size="small">详情</el-button>
                <el-button type="text" size="small">编辑</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script lang="ts" setup>

import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()

const app_conf_datum = reactive([
  {
    "name": "demo-deployment",
    "category": "Deployment",
    "order": 1,
    "description": "这个demo应用的部署组件"
  },
  {
    "name": "demo-services",
    "category": "Service",
    "order": 2,
    "description": "这个是demo应用的服务组件"
  }
])

const handleCancel = () => {
  router.back()
}

</script>
<style lang="scss">
.app-appconf-app-info {
  padding: 10px;
  background-color: white;
  border-right: rgb(228, 231, 237) 1px solid;
}
</style>