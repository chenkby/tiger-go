<template>
  <el-row class="detail-view-buttons">
    <el-col :lg="8"></el-col>
    <el-col :lg="16">
      <slot name="default">
        <el-button type="primary" plain icon="el-icon-back" @click="goBack">返回</el-button>
        <slot name="append"></slot>
      </slot>
    </el-col>

  </el-row>
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
}
body.device-mobile {
  .detail-view-buttons {
    text-align: center;
  }
}
</style>
