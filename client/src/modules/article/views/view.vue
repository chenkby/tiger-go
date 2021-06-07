<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>活动管理</el-breadcrumb-item>
    <el-breadcrumb-item>活动列表</el-breadcrumb-item>
    <el-breadcrumb-item>活动详情</el-breadcrumb-item>
  </el-breadcrumb>
  <div class="page-header">
    <div class="title">查看详情</div>
  </div>
  <el-descriptions :column="1" title="用户信息" border>
    <el-descriptions-item :span="1" label="用户名">kooriookami</el-descriptions-item>
    <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="备注">
      <el-tag size="small">学校</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="联系地址">江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item>
  </el-descriptions>

  <div class="detail-view-buttons">
    <el-row>
      <el-col>
        <el-button type="primary" plain icon="el-icon-back">返回</el-button>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  data() {
    return {

    }
  },
  computed: {
    /**
     * 列数，如果手机端，只使用1列，否则使用默认的3列
     */
    column() {
      if (this.$device === 'mobile') {
        return 1
      }
      return 3
    }
  }
}
</script>

<style lang="scss">
.detail-view-buttons {
  padding: 30px 0 30px 0;
  .el-col {
    text-align: center;
  }
}
.el-descriptions__label {
  text-align: right !important;
  &.is-bordered-label {
    font-weight: normal;
  }
}
</style>
