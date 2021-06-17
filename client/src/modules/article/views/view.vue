<template>

  <page-header title="查看详情"></page-header>
  <descriptions title="标题" extra="right">
    <el-descriptions-item label="用户名">aaa</el-descriptions-item>
    <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="备注">
      <el-tag size="small">学校</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="联系地址">江苏省苏州市吴中区吴中大道 1188 号</el-descriptions-item>
  </descriptions>
  <descriptions title="标题">
    <template #title>title1</template>
    <template #extra>
      <el-button type="primary" size="small">清除验证信息</el-button>
    </template>
    <el-descriptions-item :span="1" label="用户名">kooriookami</el-descriptions-item>
    <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="居住地">苏州市</el-descriptions-item>
    <el-descriptions-item label="备注">
      <el-tag size="small">学校</el-tag>
    </el-descriptions-item>
    <el-descriptions-item label="联系地址">江苏省苏州市深圳市龙华新区民治街道办写字楼吴中区吴中大道 1188 号</el-descriptions-item>
  </descriptions>

  <view-buttons></view-buttons>
</template>
<script>
import viewMixin from '@/common/mixins/view'
import { ref } from '@vue/reactivity'

export default {
  setup() {
    const loading = ref(true)
    setTimeout(() => {
      loading.value = false
    }, 100)

    return {
      loading
    }
  },
  mixins: [viewMixin]
}
</script>
