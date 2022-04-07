<template>
  <div>
    <el-card>
      <el-page-header :icon="ArrowLeft" content="发布任务编排"/>
      <br/>
      <el-form :model="deploy_task" label-width="120px">
        <el-form-item label="任务名称">
          <el-input v-model="deploy_task.name"></el-input>
        </el-form-item>
        <el-form-item label="发布环境">
          <el-select v-model="deploy_task.env_id">
            <el-option label="生产环境" value="prod"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="发布应用">
          <el-select v-model="deploy_task.app_id">
            <el-option label="应用1" value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="deploy_task.desciption" type="textarea"></el-input>
        </el-form-item>

      </el-form>
      <el-divider></el-divider>
      <el-row>
        <el-col :span="4" class="app-deploy-edit-left-panel">
          <el-scrollbar>
            <el-button type="default" style="width: 100%;">增加组件</el-button>
            <el-divider></el-divider>
            <div class="app-deploy-comp-list">
              <div class="app-deloy-comp-list-item">
                <div>应用1部署组件</div>
                <el-icon>
                  <ele-Remove/>
                </el-icon>
              </div>
            </div>
            <div class="app-deploy-comp-list">
              <div class="app-deloy-comp-list-item selected">
                <div>应用2部署组件</div>
                <el-icon>
                  <ele-Remove/>
                </el-icon>
              </div>
            </div>
          </el-scrollbar>
        </el-col>
        <el-col :span="20" class="app-deploy-container">
          <el-row>
            <el-col :span="11">
              <div class="app-deploy-comp-item-form">
                <h3 style="margin-left: 60px;margin-bottom: 10px;">原组件配置</h3>
                <el-form :model="deploy_origin_app_comp" label-width="120px">
                  <el-form-item label="组件名">
                    <el-input v-model="deploy_origin_app_comp.name"></el-input>
                  </el-form-item>
                  <el-form-item label="组件版本">
                    <el-input v-model="deploy_origin_app_comp.version"></el-input>
                  </el-form-item>
                  <el-form-item label="制品">
                    <el-input v-model="deploy_origin_app_comp.artifact"></el-input>
                  </el-form-item>
                  <el-form-item label="制品版本">
                    <el-input v-model="deploy_origin_app_comp.artifact_version"></el-input>
                  </el-form-item>
                  <el-form-item label="实例数量">
                    <el-input-number v-model="deploy_origin_app_comp.ins_num"></el-input-number>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="text">高级配置</el-button>
                  </el-form-item>
                </el-form>
              </div>
            </el-col>
            <el-col :span="2">
              <div style="height:100%;display:flex;flex-direction:column;justify-content: center;">
                <el-icon style="align-self: center">
                  <ele-DArrowRight/>
                </el-icon>
              </div>
            </el-col>
            <el-col :span="11">
              <h3 style="margin-left: 60px;margin-bottom: 10px;">灰度组件配置</h3>

              <el-form :model="deploy_new_app_comp" label-width="120px">
                <el-form-item label="组件名">
                  <el-input v-model="deploy_new_app_comp.name"></el-input>
                </el-form-item>
                <el-form-item label="组件版本">
                  <el-input v-model="deploy_new_app_comp.version"></el-input>
                </el-form-item>
                <el-form-item label="制品">
                  <el-input v-model="deploy_new_app_comp.artifact"></el-input>
                </el-form-item>
                <el-form-item label="制品版本">
                  <el-input v-model="deploy_new_app_comp.artifact_version"></el-input>
                </el-form-item>
                <el-form-item label="实例数量">
                  <el-input-number v-model="deploy_new_app_comp.ins_num"></el-input-number>
                </el-form-item>
                <el-form-item>
                  <el-button type="text">高级配置</el-button>
                </el-form-item>
              </el-form>
            </el-col>
          </el-row>

        </el-col>
      </el-row>
      <el-divider></el-divider>
      <div style="display: flex; justify-content: center;">
        <el-button type="primary">发布</el-button>
        <el-button>取消</el-button>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts">

import {defineComponent, reactive, toRefs} from "vue";

export default defineComponent({
  name: "app_comp_deploy_edit",
  setup() {

    const states = reactive({
      deploy_task: {
        name: "",
        app_id: "",
        env_id: "",
        desciption: "",
      },
      deploy_origin_app_comp: {
        name: "aa",
        app_id: "1",
        artifact: "xxx",
        artifact_version: "v1",
        version: "v1",
        ins_num: 1
      },
      deploy_new_app_comp: {
        name: "aa",
        app_id: "1",
        artifact: "xxx",
        artifact_version: "v1",
        version: "v1",
        ins_num: 3,
      }
    })

    return {
      ...toRefs(states)
    }
  }
})
</script>
<style lang="scss">
.app-deploy-edit-left-panel {
  border-right: #ccc solid 1px;
  padding: 0px 10px
}

.app-deploy-comp-list {
  .app-deloy-comp-list-item {
    padding: 10px 10px;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    border-left: transparent solid 3px;
    margin-bottom: 5px;

    .el-icon {
      margin-top: 4px;
      justify-self: flex-end;
      font-size: 14px;
      line-height: 14px;
    }

    &:hover, &.selected {
      border-left-color: #3eaadc;
      background-color: #ecf5ff;
    }
  }
}

.app-deploy-container {
  .app-deploy-comp-item-form {

  }
}
</style>