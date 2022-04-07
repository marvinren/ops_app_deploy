<template>
  <div class="app-appconf-container">
    <el-card shadow="hover">
      <div class="app-appconf-search mb15" >
        <span style="font-size:18px;line-height: 32px;">
          应用组件
        </span>
        <div style="float:right">
          <el-button size="default" type="default" class="ml10" @click="handleCancel">
            取消
          </el-button>
          <el-button size="default" type="primary" class="ml10">
            保存
          </el-button>
        </div>
      </div>
      <el-divider />
      <el-form :model="app_component" label-width="120px">
        <el-form-item label="组件名称">
          <el-input v-model="app_component.name"/>
        </el-form-item>
        <el-form-item label="组件类型">
          <el-select v-model="app_component.category">
            <el-option label="spring cloud 微服务应用" value="spring_cloud_micro"></el-option>
            <el-option label="k8s ingress" value="ingress"></el-option>
            <el-option label="istio 应用" value="istio"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="部署方式">
          <el-select v-model="app_component.deploy_type">
            <el-option label="k8s编排部署" value="k8s"></el-option>
            <el-option label="部署流水线" value="cdflow"></el-option>
            <el-option label="ansible脚本" value="ansible"></el-option>
            <el-option label="shell脚本" value="shell"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="所属应用">
          <el-select v-model="app_component.app_id">
            <el-option v-for="(app, index) in apps" :key="index" :label="app.name" :value="app.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="对应制品">
          <el-select v-model="app_component.artifact">
            <el-option v-for="(arti, index) in artifacts" :key="index" :label="arti.name" :value="arti.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="默认版本">
          <el-input v-model="app_component.version"></el-input>
        </el-form-item>
        <el-form-item label="构建流水线">
          <el-select v-model="app_component.ci_flow_id">
            <el-option v-for="(flow, index) in flows" :key="index" :label="flow.name" :value="flow.value"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="交付流水线">
          <el-select v-model="app_component.cd_flow_id">
            <el-option v-for="(flow, index) in flows" :key="index" :label="flow.name" :value="flow.value"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="ansible/shell脚本路径">
          <el-input v-model="app_component.script_path"></el-input>
        </el-form-item>
      </el-form>

    </el-card>
  </div>
</template>
<script lang="ts" setup>

import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()

const app_component = reactive({
  "id": "",
  "name": "",
  "category": "",
  "app_id": "1",
  "artifact": "",
  "deploy_type": "",
  "version": "",
  "ci_flow_id": "",
  "cd_flow_id": "",
  "script_path": ""
})

const apps = reactive([
  {"id": 1, "name": "应用1"},
  {"id": 2, "name": "应用2"}
])

const flows = reactive([
  {"id": 1, "name": "default-springcloud-ci-flow"},
  {"id": 2, "name": "default-springcloud-cd-flow"},
  {"id": 3, "name": "istio-ci-flow"},
  {"id": 4, "name": "istio-cd-flow"},
])
const artifacts = reactive([
  {"id": 1, "name": "nginx"},
  {"id": 2, "name": "demo-app"}
])

const handleCancel = () => {
  router.back()
}

</script>