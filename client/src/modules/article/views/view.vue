<template>
  <el-skeleton :rows="10" animated :loading="loading" :throttle="500">
    <template #default>
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>活动管理</el-breadcrumb-item>
        <el-breadcrumb-item>活动列表</el-breadcrumb-item>
        <el-breadcrumb-item>活动详情</el-breadcrumb-item>
      </el-breadcrumb>
      <div class="page-header">
        <div class="title">查看详情</div>
      </div>
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
      <descriptions title="标题" extra="right">
        <template #title>title1</template>
        <template #extra>extra2</template>
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

  </el-skeleton>
</template>
<script>
import Descriptions from '@/components/descriptions/Descriptions.vue'
import ViewButtons from '@/components/ViewButtons.vue'
import { ref } from '@vue/reactivity'

export default {
  components: { Descriptions, ViewButtons },
  setup() {
    const loading = ref(true)
    setTimeout(() => {
      loading.value = false
    }, 700)

    return {
      loading
    }
  }
}
</script>
