<template>
  <div>
    <el-card shadow="hover">
      <div class="system-menu-search mb15">
        <el-input size="default" placeholder="请输入应用组件" style="max-width: 180px"></el-input>
        <el-button size="default" type="primary" class="ml10">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
      </div>
      <el-table :data="deploy_tasks">
        <el-table-column prop="name" label="发布任务名称" width="180"></el-table-column>
        <el-table-column prop="category_name" label="类型" width="180"></el-table-column>
        <el-table-column prop="func" label="发布方式" width="180"></el-table-column>
        <el-table-column prop="start_time" label="发布开始时间" width="180"></el-table-column>
        <el-table-column prop="status" label="发布状态" width="180"></el-table-column>
        <el-table-column prop="description" label="发布描述" width="180"></el-table-column>

        <el-table-column fixed="right" label="操作" width="180">
          <template #default>
            <el-button type="text" size="small" @click="view_task_detail">任务详情</el-button>
            <el-button type="text" size="small">组件实例</el-button>
          </template>
        </el-table-column>

      </el-table>

      <el-pagination
          class="mt15"
          :pager-count="5"
          :page-sizes="[10, 20, 30]"
          v-model:current-page="tableData.param.pageNum"
          background
          v-model:page-size="tableData.param.pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="tableData.total"
      >
      </el-pagination>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const tableData = reactive({
  total: 220,
  param: {
    pageNum: 1,
    pagesize: 22
  }
})

const deploy_tasks = reactive([
  {
    name: "组件1-部署", category_name: "组件部署", func: "k8s集群部署",
    start_time: "2022-01-02 00:11:11", status: "部署中", description: "组件部署"
  },
  {
    name: "组件1-部署", category_name: "组件部署", func: "k8s集群部署",
    start_time: "2022-01-02 00:11:11", status: "部署中", description: "组件部署"
  },
  {
    name: "组件1-部署", category_name: "组件部署", func: "k8s集群部署",
    start_time: "2022-01-02 00:11:11", status: "部署中", description: "组件部署"
  },
  {
    name: "组件1-部署", category_name: "组件部署", func: "k8s集群部署",
    start_time: "2022-01-02 00:11:11", status: "部署中", description: "组件部署"
  }

])

const view_task_detail = () => {
  router.push({"path": "/app/appdeploy/view"})
}
</script>