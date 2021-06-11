<template>
  <div class="layout-sidebar">
    <div class="sidebar-menu">
      <nav-menu :is-collapse="isCollapse"></nav-menu>
    </div>

    <div class="layout-sidebar__collapse" @click="toggleSidebar">
      <i class="el-icon-s-fold"></i>
    </div>

  </div>
</template>

<script>
import NavMenu from '@/components/NavMenu.vue'

export default {
  components: { NavMenu },
  data() {
    return {
      isCollapse: false
    }
  },
  methods: {
    toggleSidebar() {
      this.isCollapse = !this.isCollapse
      this.$emit('change', this.isCollapse)
    }
  }
}
</script>

<style lang="scss">
.layout-sidebar {
  background-color: #f9fafe;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  width: $--sidebar-width;
  overflow-x: hidden;
  transition: width ease-in-out 0.2s;
}
// 最小化侧边栏
.layout-main.collapse {
  .layout-sidebar {
    width: 50px;
  }
  .layout-sidebar__collapse {
    left: 50px;
  }
}

.layout-sidebar__collapse {
  position: fixed;
  top: calc(50% - 32px);
  left: $--sidebar-width;
  z-index: 9999;
  padding: 0;
  border-width: 9px 0 9px 20px;
  border-color: transparent transparent transparent #ebebeb;
  background: none;
  width: 0;
  height: 64px;
  line-height: 46px;
  color: #000;
  border-style: solid;
  display: flex;
  align-items: center;
  cursor: pointer;
  transition: left ease-in-out 0.2s;
  user-select: none;
  i {
    position: relative;
    left: -18px;
  }
}
.el-submenu [class^='el-icon-'],
.el-menu-item [class^='el-icon-'] {
  width: 20px;
  margin-right: 15px;
}
.el-submenu__title,
.el-menu-item {
  line-height: normal;
  display: flex;
  align-items: center;
  flex-wrap: nowrap;
}
.nav-menu > .el-submenu > .el-submenu__title,
.nav-menu > .el-menu-item {
  padding-left: 15px !important;
}
</style>
