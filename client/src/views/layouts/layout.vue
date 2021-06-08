<template>
  <div class="layout-container">
    <layout-header></layout-header>
    <div class="layout-main" :class="{collapse:isCollapse}">
      <el-scrollbar :native="scrollbarNative">
        <div class="layout-main__body-container">
          <layout-sidebar @change="onCollapseChange"></layout-sidebar>
          <div class="layout-page-container">

            <router-view></router-view>

          </div>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<script>
import LayoutHeader from './header.vue'
import LayoutSidebar from './sidebar.vue'

export default {
  components: { LayoutHeader, LayoutSidebar },
  data() {
    return {
      isCollapse: false
    }
  },
  methods: {
    /**
     * 收缩侧边栏事件
     */
    onCollapseChange(isCollapse) {
      this.isCollapse = isCollapse
    }
  },
  computed: {
    /**
     * 如果是手机端，使用原生的滚动条样式(更细)
     */
    scrollbarNative() {
      return this.$isMobile
    }
  }
}
</script>
<style lang="scss">
</style>
