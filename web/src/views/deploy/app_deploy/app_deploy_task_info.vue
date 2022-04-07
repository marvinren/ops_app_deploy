<template>
  <div>
    <el-card>
      <div>
        <el-button>查看应用组件</el-button>
        <el-button>查看部署日志</el-button>
        <el-button>黑屏操作</el-button>
        <el-button>组件运维</el-button>
      </div>
      <el-divider></el-divider>
      <el-timeline>
        <el-timeline-item center timestamp="2018/4/2 20:46" placement="top">
          发布开始
        </el-timeline-item>

        <el-timeline-item center timestamp="2018/4/12 20:46" placement="top">
          <el-card shadow="never">
            <h4>发布校验</h4>
            <p>制品版本校验 2018/4/12 20:46</p>
            <p>制品质量门禁校验 2018/4/12 20:46</p>
          </el-card>
        </el-timeline-item>
        <el-timeline-item timestamp="2018/4/3 20:46" placement="top">
          <el-card shadow="never">
            <h4>组件发布</h4>
            <p>发布执行 2018/4/12 20:46</p>
            <p>发布完成 2018/4/12 20:46</p>
            <p>发布后维持30s 2018/4/12 20:46</p>
          </el-card>
        </el-timeline-item>

        <el-timeline-item timestamp="2018/4/2 20:46" placement="top">
          发布结束
        </el-timeline-item>
      </el-timeline>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
</script>