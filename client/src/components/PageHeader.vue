<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>活动管理</el-breadcrumb-item>
    <el-breadcrumb-item>活动列表</el-breadcrumb-item>
    <el-breadcrumb-item>活动详情</el-breadcrumb-item>
  </el-breadcrumb>
  <div class="page-header">
    <div class="page-header__left">
      <slot name="default">
        <div class="title">{{title}}</div>
      </slot>
    </div>
    <div class="page-header__right">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  setup() {

  },
  props: {
    title: {
      type: String,
      default: '页面标题'
    }
  }
}
</script>

<style lang="scss">
.page-header {
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
  .title {
    font-size: 24px;
    font-weight: 400;
    color: #111;
  }
}
body.device-mobile {
  .page-header {
    display: none;
  }
}
</style>