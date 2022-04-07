<template>
  <div>

    <el-card>
      <el-page-header :icon="ArrowLeft" content="发布任务"/>

      <br/>
      <div class="system-menu-search mb15">
        <el-input placeholder="请输入任务名称" style="max-width: 180px"></el-input>
        <el-button type="primary" class="ml10">
          <el-icon>
            <ele-Search/>
          </el-icon>
          查询
        </el-button>
        <el-button type="default" class="ml15" @click="onEdit">
          <el-icon>
            <ele-Plus/>
          </el-icon>
          新建发布任务
        </el-button>
      </div>

      <el-tabs type="border-card" v-model="activeTabName">
        <el-tab-pane label="发布任务" name="tasks">
          <el-table :data="tasks">
            <el-table-column prop="name" label="任务名称" width="120"></el-table-column>
            <el-table-column prop="app_name" label="所属应用" width="120"></el-table-column>
            <el-table-column prop="env_name" label="部署环境" width="120"></el-table-column>
            <el-table-column prop="comp_num" label="组件数量" width="120"></el-table-column>
            <el-table-column prop="status" label="状态" width="120"></el-table-column>
            <el-table-column prop="create_at" label="创建时间" width="120"></el-table-column>
            <el-table-column prop="create_by" label="创建人" width="120"></el-table-column>
            <el-table-column prop="description" label="描述"></el-table-column>
            <el-table-column fixed="right" label="操作">
              <template #default>
                <el-button type="text" size="small">编辑</el-button>
                <el-button type="text" size="small">启动</el-button>
                <el-button type="text" size="small">删除</el-button>
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
        </el-tab-pane>
        <el-tab-pane label="历史任务" name="history_tasks">
          <div>...</div>
        </el-tab-pane>
      </el-tabs>

    </el-card>

  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, ref, toRefs} from "vue";
import {useRouter} from "vue-router";

export default defineComponent({
  name: "app_grey_deploy_index",
  setup() {
    const states = reactive({
      activeTabName: "tasks",
      tasks: [{
        name: "任务1",
        app_name: "应用1",
        env_name: "生产环境",
        status: "运行中",
        comp_num: 1,
        create_at: "2022-03-01 10:00:00",
        create_by: "admin",
        description: ""
      }, {
        name: "任务1",
        app_name: "应用1",
        env_name: "生产环境",
        status: "运行中",
        comp_num: 1,
        create_at: "2022-03-01 10:00:00",
        create_by: "admin",
        description: ""
      }],
      tableData: {
        total: 220,
        param: {
          pageNum: 1,
          pagesize: 22
        }
      }
    })

    const router = useRouter()
    const onEdit = ()=>{
      router.push({'path': '/app_grey_deploy/normal_edit'})
    }
    return {
      onEdit,
      ...toRefs(states),

    }
  }
})
</script>