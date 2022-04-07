<template>
  <el-form :model="environment" label-width="120px">
    <el-form-item label="名称">
      <el-input v-model="environment.name" />
    </el-form-item>
    <el-form-item label="类型">
      <el-select v-model="environment.category">
        <el-option label="生产" value="prod"></el-option>
        <el-option label="准生产" value="uat"></el-option>
        <el-option label="测试" value="test"></el-option>
        <el-option label="开发" value="dev"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="描述">
      <el-input v-model="environment.description" type="textarea"/>
    </el-form-item>
    <el-form-item label="所属云环境">
      <el-select v-model="environment.cloud_id">
        <el-option label="阿里云" value="ali1"></el-option>
        <el-option label="华为云" value="huawei1"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="基础资源">
      <div style="background-color: #a5a5a5;width:100%;padding:5px;">
        <el-row gutter="10">
          <el-col :span="8">
            <el-card class="box-card">
              <el-descriptions style="width:100%" column="1">
                <el-descriptions-item label="集群" label-align="left">ccs.app.demo</el-descriptions-item>
                <el-descriptions-item label="规格">ecs.small.10g</el-descriptions-item>
                <el-descriptions-item label="状态">正常</el-descriptions-item>
              </el-descriptions>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card class="box-card">
              <el-descriptions style="width:100%" column="1">
                <el-descriptions-item label="集群" label-align="left">k8s.demo.prod</el-descriptions-item>
                <el-descriptions-item label="规格">100core, 4095GB</el-descriptions-item>
                <el-descriptions-item label="状态">正常</el-descriptions-item>
              </el-descriptions>
            </el-card>
          </el-col>
        </el-row>

      </div>
      <el-row>
        <el-col :span="24">
          <el-button type="default" style="margin-top:5px;">新增基础资源</el-button>
        </el-col>
      </el-row>
    </el-form-item>
  </el-form>
</template>
<script>

export default {
  data: () => ({
    environment: {
      "id": 1,
      "name": "",
      "category": "",
      "description": "",
      "cloud_id": "",
      "cluster": []
    }
  })
}
</script>