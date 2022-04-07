<template>
  <div class="app-appconf-container">
    <el-card shadow="hover">
      <div class="app-appconf-search mb15" >
        <span style="font-size:18px;line-height: 32px;">
          应用组件变量
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
        <el-col :span="6">
          <div class="app-appconf-app-info">
            <div class="appconf-group-header">
              <span style="line-height: 40px;font-size: 16px;font-weight: 600;">变量组</span>

              <div style="float:right">
                <el-button type="text">新增</el-button>
              </div>
            </div>

            <div class="appconf-group-list">
              <div class="appconf-group-item" :class="{selected: item.id === selected_item_id}"
                   v-for="(item, index) in conf_group" :key="index">
                <span>{{ item.name }}</span><span style="float:right;color: dodgerblue;">{{ item.env }}</span>
              </div>
            </div>
          </div>
        </el-col>

        <el-col :span="18">
          <el-table :data="selected_group_variables" style="width: 100%">
            <el-table-column label="变量名" width="180">
              <template #default="scope">
                <el-input v-model="scope.row.name"></el-input>
              </template>
            </el-table-column>
            <el-table-column label="变量值" width="200">
              <template #default="scope">
                <el-input v-model="scope.row.v"></el-input>
              </template>
            </el-table-column>
            <el-table-column label="描述">
              <template #default="scope">
                <el-input v-model="scope.row.desc"></el-input>
              </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" width="120">
              <template #default="scope">
                <el-button type="text" size="small">
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
          <div style="width: 100%;padding:5px 15px">
            <el-button class="mt-4" style="width: 100%;">添加</el-button>
          </div>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const selected_item_id = ref(1)
const conf_group = reactive([
  {'id': 1, 'name': '准生产变量组', 'env': 'uat'},
  {'id': 2, 'name': '生产变量组', 'env': 'prod'},
  {'id': 3, 'name': '开发变量组', 'env': 'dev'}
])
const selected_group_variables = reactive([
  {"name": "namespace", "v": "3", "desc": "hello"},
  {"name": "namespace", "v": "3", "desc": "hello"},
  {"name": "namespace", "v": "3", "desc": "hello"},
  {"name": "namespace", "v": "3", "desc": "hello"},
])

const handleCancel = ()=>{
  router.back()
}

</script>
<style lang="scss">


.app-appconf-app-info {
  padding: 10px;
  background-color: white;
  border-right: rgb(228, 231, 237) 1px solid;
  height: 100%;
}

.appconf-group-header {
  padding: 5px 5px;
}

.appconf-group-list {
  .appconf-group-item {
    padding: 8px 5px;
    margin-bottom: 10px;
    font-size: 14px;
    line-height: 14px;

    &:hover {
      background-color: #ecf5ff;
    }

    &.selected {
      background-color: #ecf5ff;
    }
  }
}
</style>