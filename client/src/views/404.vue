<template>
  <el-result :title="title" :subTitle="subTitle">
    <template #icon>
      <el-image :src="import('../assets/images/404.png')"></el-image>
    </template>
    <template #extra>
      <el-button type="primary" size="medium" @click="goBack">返回</el-button>
    </template>
  </el-result>
</template>

<script>
export default {
  props: {
    title: {
      type: [String, Number],
      default: 404
    },
    subTitle: {
      type: String,
      default: '您访问的页面不存在'
    }
  },
  methods: {
    goBack() {
      this.$router.back()
    }
  }
}
</script>
