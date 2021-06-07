<template>
  <div class="detail-view-buttons">
    <slot name="default">
      <el-button type="primary" plain icon="el-icon-back" @click="goBack">返回</el-button>
      <slot name="append"></slot>
    </slot>

  </div>
</template>

<script>
/**
 * slots:
 * default:重写内容
 * append:在后退按钮后面添加内容
 */
export default {
  methods: {
    goBack() {
      this.$router.back()
    }
  }
}
</script>
<style lang='scss'>
.detail-view-buttons {
  padding: 30px 0 30px 0;
  text-align: center;
}
</style>