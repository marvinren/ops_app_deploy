<template>
  <div style="height:100%">
    <el-card shadow="hover">
      <div class="app-appconf-search mb15" >
        <span style="font-size:18px;line-height: 32px;">
          应用组件编排
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
      <el-row>
        <el-col :span="14" class="app-appconf-code-editor">
          <el-tabs v-model="activeName" editable>
            <el-tab-pane v-for="(c, idx) in app_confs" :key="idx" :label="c.confName">
              <app-conf-yaml-editor :conf-name="c.confName" :conf-script="c.confScript"></app-conf-yaml-editor>
            </el-tab-pane>
          </el-tabs>
        </el-col>
        <el-col :span="10" class="app-appconf-variable-panel">
          <h3>预置变量</h3>
          <el-table :data="presetValues" border style="width: 100%">
            <el-table-column prop="v_name" label="名称"/>
            <el-table-column prop="v_category" label="类型"/>
            <el-table-column prop="v_source" label="值来源"/>
            <el-table-column prop="v_value" label="值"/>
          </el-table>
          <br/>
          <h3>自定义变量</h3>
          <el-table :data="customValues" border style="width: 100%">
            <el-table-column prop="v_name" label="名称"/>
            <el-table-column prop="v_category" label="类型"/>
            <el-table-column prop="v_desc" label="描述"/>
            <el-table-column prop="v_value" label="值"/>
          </el-table>

        </el-col>
      </el-row>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
import AppConfYamlEditor from './component/app_conf_yaml_editor.vue'
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const activateName = ref("demo_deployment")
const app_confs = reactive([{
  "confName": "demo_deployment", "confScript": `apiVersion: extensions/v1beta1   #当前格式的版本
kind: Deployment                 #当前创建资源的类型， 当前类型是Deployment
metadata:                        #当前资源的元数据
  name: nginx-test               #当前资源的名字 是元数据必须的项
spec:                            #是当前Deployment的规格说明
  replicas:                      #指当前创建的副本数量 默认不填 默认值就为‘1’
  template:                      #定义pod的模板
    metadata:                    #当前pod的元数据
      labels:                    #至少顶一个labels标签，可任意创建一个 key:value
        app: web_server
    spec:                        #当前pod的规格说明
      containers:                #容器
      - name: nginx              #是容器的名字容器名字是必须填写的
        image: nginx:latest      #镜像 镜像的名字和版本`
}, {
  "confName": "demo_service", "confScript": `apiVersion : v1 #版本v1
kind : Pod  #类型是pod
metadata:  #对象需要的属性值
  name: nginx  #name为nginx
spec:  #定义容器参数
  containers:
    - name: nginx
      image: nginx #镜像为nginx
      ports:
      - containerport: 80 #容器本地端口`
}
])
const presetValues = reactive( [
  {"v_name": "namespace", "v_category": "string", "v_source": "const", "v_value": "default"},
  {"v_name": "relicas", "v_category": "int", "v_source": "const", "v_value": "2"},
  {"v_name": "image", "v_category": "string", "v_source": "const", "v_value": "nginx"},
  {"v_name": "envid", "v_category": "string", "v_source": "const", "v_value": "dev"},
])

const customValues = reactive([
  {"v_name": "namespace", "v_category": "string", "v_value": "default", "v_desc": "命名空间"},
])

const handleCancel = () => {
  router.back()
}

</script>
<style lang="scss">
.app-appconf-code-editor {
  border-right: rgb(228, 231, 237) 1px solid;

  .el-tabs__new-tab {
    margin-right: 10px;
  }
}

.app-appconf-variable-panel {
  padding: 10px;
}
</style>