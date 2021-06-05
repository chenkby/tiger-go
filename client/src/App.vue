<template>

  <div class="layout-container">
    <el-header height="50px" class="layout-header">
      <div class="app-logo">
        <img class="icon" src="https://ergev2.tuxiaobei.com/manage/img/logo.93fa7efe.png" />
        <div class="name">兔小贝儿歌</div>
      </div>

      <div class="layout-header__tools">
        <el-space size="large">
          <div class="item">
            <el-link :underline="false">消息</el-link>
          </div>
          <div class="item">
            <el-link :underline="false">通知</el-link>
          </div>
        </el-space>
      </div>

      <div class="layout-header__user">
        <el-dropdown>
          <div class="el-dropdown-link">
            <el-avatar :size="28" src="https://ergev2.tuxiaobei.com/manage/img/logo.93fa7efe.png"></el-avatar>
          </div>
          <template #dropdown>
            <el-dropdown-menu class="layout-header__popup">
              <div class="layout-header__user-info">
                <el-avatar class="user-avatar" :size="28" src="https://ergev2.tuxiaobei.com/manage/img/logo.93fa7efe.png"></el-avatar>
                <span class="user-nickname">超级管理员</span>
              </div>
              <el-dropdown-item icon="el-icon-lock">修改密码</el-dropdown-item>
              <el-dropdown-item icon="el-icon-aim">狮子头</el-dropdown-item>
              <el-dropdown-item icon="el-icon-apple">螺蛳粉</el-dropdown-item>
              <el-dropdown-item divided icon="el-icon-switch-button" class="user-logout">退出登录</el-dropdown-item>

            </el-dropdown-menu>

          </template>
        </el-dropdown>
      </div>
    </el-header>
    <div class="layout-main">
      <el-scrollbar>
        <div class="layout-main__body-container">
          <el-scrollbar class="layout-sidebar">

            <div class="sidebar-menu">
              <el-menu default-active="2" class="el-menu-vertical-demo" background-color="#f5f5f5" text-color="#333" active-text-color="#000">
                <el-submenu index="1">
                  <template #title>
                    <i class="el-icon-location"></i>
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
                  <template #title>导航二</template>
                </el-menu-item>
                <el-menu-item index="3" disabled>
                  <i class="el-icon-document"></i>
                  <template #title>导航三</template>
                </el-menu-item>
                <el-menu-item index="4">
                  <i class="el-icon-setting"></i>
                  <template #title>导航四</template>
                </el-menu-item>
              </el-menu>
            </div>

            <div class="sidebar-collapse">
              <i class="el-icon-s-fold"></i>
            </div>

          </el-scrollbar>
          <div class="layout-page-container">

            <router-view></router-view>

          </div>
        </div>
      </el-scrollbar>
    </div>

  </div>

</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  data() {
    return {

    }
  },
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
html,
body {
  margin: 0;

  width: 100%;
  height: 100%;
  overflow: hidden;
}
body {
  padding: 50px 0 0 0;
  height: calc(100% - 50px);
}

#app {
  height: 100vh;
}
.layout-container {
  width: 100%;
  height: 100%;
  overflow-y: auto;
}
.layout-header {
  z-index: 9;
  box-shadow: 0 2px 4px 0 var(--cb-color-shadow, rgba(0, 0, 0, 0.16));
  display: flex;
  align-items: center;
  justify-content: flex-end;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background-color: #fff;
  .app-logo {
    margin-right: auto;
  }
  .layout-header__tools {
    margin-right: 20px;
    .el-link {
      font-weight: normal;
    }
  }
  .layout-header__user {
    cursor: pointer;
    .el-dropdown-link {
      width: 40px;
      height: 50px;
      display: flex;
      align-items: center;
    }
  }
}
.layout-header__popup {
  width: 220px;
  padding: 5px 0;
  .layout-header__user-info {
    padding: 10px 15px 15px 15px;
    border-bottom: 1px solid #efefef;
    display: flex;
    align-items: center;
    margin-bottom: 5px;
    .user-nickname {
      margin-left: 10px;
      font-size: 14px;
    }
  }
  .user-logout {
    text-align: center;
  }
}
.layout-main {
  padding: 0;
  margin-top: 0;
  height: 100%;
  min-height: auto;
  box-sizing: border-box;
}
.layout-sidebar {
  background-color: #f5f5f5;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  width: 240px;
  transition: max-width ease-in-out 0.3s;
}
.layout-main__body-container {
  height: 100%;
  box-sizing: border-box;
}
.layout-page-container {
  padding-bottom: 100px;
  padding: 20px 30px 100px 270px;
}
.app-logo {
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  .icon {
    width: 30px;
    height: 30px;
  }
  .name {
    color: #2c3e50;
    font-size: 20px;
    font-weight: normal;
    margin-left: 10px;
  }
}
.sidebar-menu {
  .el-menu {
    border-right: none;
  }
}
.sidebar-collapse {
  display: none;
  position: absolute;
  top: calc(50% - 32px);
  right: -18px;
  z-index: 9999;
  padding: 0;
  border-width: 9px 0 9px 20px;
  border-color: transparent transparent transparent #ebebeb;
  background: none;
  width: 0;
  height: 64px;
  line-height: 46px;
  color: #000;
  font-size: 12px;
  border-style: solid;
  i {
    position: relative;
    left: -15px;
  }
}

.layout-container--main {
  margin-left: 250px;
}
.layout__main {
  background-color: #f5f5f5;
  padding: 10px !important;
}
.el-header {
  display: flex;
  align-items: center;
  .collapse {
    cursor: pointer;
    width: 40px;
    height: 40px;
    justify-content: center;
    display: flex;
    align-items: center;
    i {
      font-size: 22px;
      color: #666;
    }
  }
}

/** breadcrumb **/
.el-breadcrumb {
  font-size: 12px !important;
  margin-bottom: 15px;
}
.el-breadcrumb__inner {
  color: #888;
}
.el-breadcrumb__inner.is-link,
.el-breadcrumb__inner a {
  font-weight: normal;
  color: #333;
}

.page-header {
  height: 40px;
  display: flex;
  align-items: center;
  margin-bottom: 15px;
  .title {
    font-size: 24px;
    font-weight: 400;
    color: #111;
  }
}
</style>
