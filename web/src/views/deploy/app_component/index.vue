<template>
  <div class="app-mgr-container">
    <el-card shadow="hover">
      <div class="system-menu-search mb15">
        <el-input size="default" placeholder="请输入应用组件" style="max-width: 180px"> </el-input>
        <el-button size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search />
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="primary" class="ml10">批量删除</el-button>
        <el-button size="default" type="primary" class="ml10">更新制品版本</el-button>
        <el-button size="default" type="success" class="ml10" @click="edit_comp">
          <el-icon>
            <ele-FolderAdd />
          </el-icon>
          新增应用组件
        </el-button>
      </div>
      <el-table :data="app_components">
        <el-table-column prop="name" label="应用组件名称" width="180"></el-table-column>
        <el-table-column prop="app_name" label="所属应用" width="180"></el-table-column>
        <el-table-column prop="category" label="组件类型" width="180"></el-table-column>
        <el-table-column prop="artifact" label="对应制品" width="180"></el-table-column>
        <el-table-column prop="deploy_type" label="部署方式" width="180"></el-table-column>
        <el-table-column prop="runtime_type" label="运行方式" width="180"></el-table-column>
        <el-table-column prop="cd_flow_id" label="流水线" width="180"></el-table-column>
        <el-table-column prop="version" label="版本" width="180"></el-table-column>
        <el-table-column fixed="right" label="操作" width="180">
          <template #default>
            <el-button type="text" size="small" @click="k8s_edit">部署编排</el-button>
            <el-button type="text" size="small" @click="conf_edit">变量配置</el-button>
            <el-button type="text" size="small" @click="deploy_comp">部署</el-button>
            <el-button type="text" size="small" @click="edit_comp">编辑</el-button>
            <el-button type="text" size="small">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script lang="ts" setup>

import {reactive, ref} from "vue";
import {useRouter} from "vue-router";


const router = useRouter()
const app_components=reactive([
      {"name": "nginx前端", "app_name":"应用1", "category": "spring cloud", "artifact": "nginx",
        "deploy_type": "k8s", "cd_flow_id": "none", "version": "v3", "env": "UAT, PROD"},
    ])

const k8s_edit = () => {
  router.push({path:`/app/appcomp/k8sconf`})
}
const conf_edit = () => {
  router.push({path: '/app/appcomp/conf'})
}
const edit_comp = () => {
  router.push({path: '/app/appcomp/edit'})
}
const deploy_comp = () => {
  router.push({path: '/app/appcomp/deploy'})
}

</script>