<template>
  <keep-alive>
    <router-view v-if="$route.meta.keepAlive"></router-view>
  </keep-alive>
  <router-view v-if="!$route.meta.keepAlive"></router-view>
</template>

<script>
export default {

  data() {
    return {

    }
  },
  mounted() {
    // 给body设置一个标识设备类型的class name
    document.getElementsByTagName('body')[0].className = `device-${this.$device}`
  }

}

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/script-setup-2/active-rfcs/0000-script-setup.md
</script>

<style lang="scss">
* {
  font-family: roboto, Avenir, Helvetica, Microsoft YaHei, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -webkit-tap-highlight-color: transparent;
}
//取消ios下的输入框阴影
input[type='password'],
input[type='text'],
textarea {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  -webkit-tap-highlight-color: transparent !important;
}
html,
body {
  margin: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

#app {
  height: 100vh;
}
// 手机端隐藏
body.device-mobile {
  .layout-sidebar {
    display: none;
  }
  .layout-page-container {
    margin-left: 0;
    padding: 20px $--mobile-padding 120px $--mobile-padding;
  }
  .quick-nav {
    display: none;
  }
}
.layout-container {
  width: 100%;
  height: 100%;
  overflow-y: auto;
}

.layout-main {
  padding: 50px 0 0 0;
  height: 100%;
  min-height: auto;
  box-sizing: border-box;
}

// 最小化侧边栏
.layout-main.collapse {
  .layout-page-container {
    margin-left: 50px;
  }
}

.layout-main__body-container {
  height: 100%;
  box-sizing: border-box;
}
.layout-page-container {
  padding: 20px 30px 20px 30px;
  transition: margin ease-in-out 0.2s;
  margin-left: $--sidebar-width;
}

.sidebar-menu {
  .el-menu {
    border-right: none;
  }
}

.layout__main {
  background-color: #f5f5f5;
  padding: 10px !important;
}

/** breadcrumb **/
.el-breadcrumb {
  font-size: 12px !important;
  margin-bottom: 20px;
}
.el-breadcrumb__inner {
  color: #888;
}
.el-breadcrumb__inner.is-link,
.el-breadcrumb__inner a {
  font-weight: normal;
  color: #333;
}

// 其他的一些设置 ****/
body.device-mobile {
  // messageBox宽度 100%
  .el-message-box {
    width: 94%;
  }
}
</style>
