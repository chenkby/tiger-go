<template>
  <el-result :title="title" :subTitle="subTitle">
    <template #icon>
      <el-image :src="defaultImage"></el-image>
    </template>
    <template #extra>
      <el-button type="primary" size="medium" @click="goBack">返回</el-button>
    </template>
  </el-result>
</template>

<script>
import defaultImage from '@/assets/images/404.png'

export default {
  data() {
    return {
      defaultImage
    }
  },
  props: {
    /**
     * 主标题
     */
    title: {
      type: [String, Number],
      default: '您访问的页面不存在'
    },
    /**
     * 副标题
     */
    subTitle: {
      type: String,
      default: '请确认链接正确。如果问题仍然存在，请联系系统管理员'
    }
  },
  methods: {
    /**
     * 返回上一页
     */
    goBack() {
      this.$router.back()
    }
  }
}
</script>
