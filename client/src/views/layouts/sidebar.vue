<template>
  <el-scrollbar class="layout-sidebar">
    <div class="sidebar-menu">
      <div class="test">
        <span class="el-icon-menu"></span>
        <span class="txt">导航一</span>
        <i class="el-submenu__icon-arrow el-icon-arrow-down"></i>
      </div>
      <el-menu :collapse="isCollapse" default-active="2" class="el-menu-vertical-demo" background-color="#f5f5f5" text-color="#333" active-text-color="#000">
        <el-submenu index="1">
          <template #title>
            <i class="el-icon-menu"></i>
            <span>导航一</span>
          </template>
          <el-menu-item-group>
            <template #title>分组一</template>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="分组2">
            <el-menu-item index="1-3">选项3</el-menu-item>
          </el-menu-item-group>
          <el-submenu index="1-4">
            <template #title>选项4</template>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-submenu>
        </el-submenu>
        <el-menu-item index="2">
          <template #title>
            <i class="el-icon-menu"></i>
            <span>导航二</span>
          </template>
        </el-menu-item>
        <el-menu-item index="3" disabled>
          <i class="el-icon-document"></i>
          <template #title>导航三</template>
        </el-menu-item>
        <el-menu-item index="4">

          <template #title>
            <i class="el-icon-setting"></i>
            <span>导航四</span>
          </template>
        </el-menu-item>
      </el-menu>
    </div>

    <div class="layout-sidebar__collapse" @click="toggleSidebar">
      <i class="el-icon-s-fold"></i>
    </div>

  </el-scrollbar>
</template>

<script>
export default {
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
  background-color: #f5f5f5;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  width: 240px;
  transition: width ease-in-out 0.2s;
}

.layout-sidebar__collapse {
  position: fixed;
  top: calc(50% - 32px);
  left: 240px;
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
  i {
    position: relative;
    left: -18px;
  }
}
.el-submenu [class^='el-icon-'],
.el-menu-item [class^='el-icon-'] {
  width: auto;
  margin-right: 0;
}
.el-submenu__title,
.el-menu-item {
  line-height: normal;
  display: flex;
  align-items: center;
  padding: 0 15px;
  flex-wrap: nowrap;
  span {
    //margin-left: 15px;
  }
}
.test * {
  vertical-align: middle;
}
.test {
  display: flex;
  align-items: center;
  padding: 15px 20px 15px 24px;
  position: relative;
  box-sizing: border-box;
  height: 50px;
  line-height: normal;
  span {
    display: inline-block;
  }
  .el-icon-s-cooperation,
  .el-icon-menu {
    font-size: 18px;
    margin-right: 0;
    width: auto;
  }
  .txt {
    font-size: 14px;
  }
  .el-submenu__icon-arrow {
    position: absolute;
    top: 50%;
    right: 20px;
    margin-top: -7px;
    transition: transform 0.3s;
    font-size: 12px;
  }
}
</style>
