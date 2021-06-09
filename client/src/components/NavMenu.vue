<template>
  <el-menu :collapse="isCollapse" default-active="2" class="nav-menu" :background-color="bgColor" :text-color="textColor" :active-text-color="activeTextColor">
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

      <i class="el-icon-menu"></i>
      <span>导航二</span>

    </el-menu-item>
    <el-menu-item index="3" disabled>
      <i class="el-icon-document"></i>
      <span>导航三</span>
    </el-menu-item>
    <el-menu-item index="4">

      <template #title>
        <i class="el-icon-setting"></i>
        <span>导航四</span>
      </template>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  props: {
    isCollapse: {
      type: Boolean,
      deafult: false
    },
    theme: {
      type: String,
      default: 'light'
    }
  },
  computed: {
    bgColor() {
      return this.theme === 'light' ? '#f5f5f5' : '#2f3447'
    },
    textColor() {
      return this.theme === 'light' ? '#333' : '#fff'
    },
    activeTextColor() {
      return this.theme === 'light' ? '#000' : '#ffd04b'
    }
  }
}
</script>

<style lang="scss">
.nav-menu {
  border-right: none;
}
</style>
