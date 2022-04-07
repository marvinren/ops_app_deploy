<template>
  <div>
    <div class="app-env-header">
      <el-card class="box-card" size="small">
        <el-button type="primary" size="default" @click="openDialog">新建环境</el-button>
        <el-button size="default">新建部署</el-button>
      </el-card>
    </div>
    <el-row :gutter="30" class="app-env-body">

      <el-col :span="6" v-for="(e,index) in environments" :index="index">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <div style="margin-bottom: 8px;"><h3>{{ e.name }}</h3></div>
              <div style="margin-bottom: 8px">
                <el-tag type="danger" size="small">{{ e.category }}</el-tag>
                &nbsp;
                <el-tag type="default" size="small">{{ e.app_name }}</el-tag>
              </div>
              <div style="margin-top:5px;color: #aaa;">
                由 {{ e.createBy }} 创建于 {{ e.createAt }}
              </div>
            </div>
          </template>
          <el-descriptions style="width:100%" column="1">
            <el-descriptions-item label="环境版本" label-align="left">{{ e.version }}</el-descriptions-item>
            <el-descriptions-item label="环境状态">{{ e.status }}</el-descriptions-item>
          </el-descriptions>
          <el-button type="default">立即部署</el-button>
          <el-button type="default">环境编排</el-button>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog
        v-model="editDialogVisible"
        title="应用"
        width="60%"
        :before-close="handleClose"
    >
      <environment-edit></environment-edit>
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

<script>
import EnvironmentEdit from "./environment_edit.vue"

export default {
  components: {EnvironmentEdit},
  data: () => ({
    editDialogVisible: false,
    environments: [
      {
        "id": 1, "name": "广州主机房-生产环境", "app_name": "应用1", "category": "生产", "description": "广州三层机房华为环境",
        "createAt": "2022-10-11 20:00:00", "createBy": "admin", "version": "v1", "status": "正在运行"
      },
      {
        "id": 1, "name": "广州主机房-生产环境", "app_name": "应用1", "category": "准生产", "description": "广州三层机房华为环境",
        "createAt": "2022-10-11 20:00:00", "createBy": "admin", "version": "v1", "status": "正在运行"
      },
    ]
  }),
  methods: {
    handleClose() {
      this.editDialogVisible = false
    },
    openDialog() {
      this.editDialogVisible = true
    }
  }
}
</script>
<style lang="scss">
.app-env-header {
  margin-bottom: 5px;

  .el-card__body {
    padding: 10px;
  }
}

.app-env-body {
  margin-top: 15px;
  .el-descriptions__content:not(.is-bordered-label) {
    float: right;
  }
}

</style>
