<template>
  <div class="app-appconf-container">
    <el-card shadow="hover">
      <div class="app-appconf-search mb15">
        <span style="font-size:18px;line-height: 32px;">
          应用发布
        </span>
        <div style="float:right">
          <el-button size="default" type="default" class="ml10" @click="handleCancel">
            取消
          </el-button>
          <el-button size="default" type="primary" class="ml10" @click="view_task_detail">
            执行部署
          </el-button>
        </div>
      </div>
      <el-divider/>
      <el-form :model="app_comp_deploy_settings" label-width="120px" >
        <el-form-item label="应用">
          <el-input v-model="app_comp_deploy_settings.app_name" disabled></el-input>
        </el-form-item>
        <el-form-item label="应用组件">
          <el-input v-model="app_comp_deploy_settings.app_comp_name" disabled></el-input>
        </el-form-item>
        <el-form-item label="对应制品">
          <el-input v-model="app_comp_deploy_settings.artifact" disabled></el-input>
        </el-form-item>
        <el-form-item label="发布环境">
          <el-select v-model="app_comp_deploy_settings.env_id">
            <el-option label="生产环境" value="prod"></el-option>
            <el-option label="准生产环境" value="uat"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="版本">
          <el-input v-model="app_comp_deploy_settings.version"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="app_comp_deploy_settings.description" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="基础资源">
          <el-select v-model="app_comp_deploy_settings.cluster">
            <el-option label="集群1" value="1"></el-option>
            <el-option label="集群2" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="实例数量">
          <el-input-number v-model="app_comp_deploy_settings.instance_num"></el-input-number>
        </el-form-item>
        <el-form-item label="CPU最小配额">
          <el-input v-model="cpu_min"></el-input>
        </el-form-item>
        <el-form-item label="CPU最大配额">
          <el-input v-model="cpu_max"></el-input>
        </el-form-item>
        <el-form-item label="内存最小配额">
          <el-input v-model="mem_min"></el-input>
        </el-form-item>
        <el-form-item label="CPU最小配额">
          <el-input v-model="mem_max"></el-input>
        </el-form-item>

      </el-form>
    </el-card>
  </div>
</template>


<script lang="ts" setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const app_comp_deploy_settings = reactive({
  app_name: "应用1",
  app_comp_name: "微服务1",
  artifact: "nginx",
  env_id: "prod",
  env_name: "准生产环境1",
  version: "v1",
  description: "",
  cluster: 1,
  instance_num: 2,
  cpu_min: 0.25,
  cpu_max: 2,
  mem_min: 0.5,
  mem_max: 2
})

const view_task_detail = () => {
  router.push({"path": "/app/appdeploy/view"})
}
</script>