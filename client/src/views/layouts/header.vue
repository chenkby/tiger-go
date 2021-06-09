<template>
  <el-header height="50px" class="layout-header">
    <div class="layout-header__brand">
      <template v-if="$isMobile">
        <div class="nav-menu" @click="showMenu">
          <tiger-icon icon="menu"></tiger-icon>
        </div>
        <div class="page-title">页面标题</div>
      </template>
      <brand v-else></brand>
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

  <!-- 移动端的侧边菜单 -->
  <el-drawer title="导航菜单" :with-header="false" size="60%" v-model="menuStatus" direction="ltr" custom-class="nav-menu__drawer">
    <div class="nav-menu__brand">
      <brand color="#fff"></brand>
    </div>
    <nav-menu theme="drak"></nav-menu>
  </el-drawer>
</template>

<script>
import Brand from '@/components/Brand.vue'
import TigerIcon from '@/components/TigerIcon.vue'
import NavMenu from '@/components/NavMenu.vue'

export default {
  components: { Brand, TigerIcon, NavMenu },
  data() {
    return {
      menuStatus: false
    }
  },
  methods: {
    showMenu() {
      this.menuStatus = true
    }
  }
}
</script>

<style lang="scss">
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

  .layout-header__tools {
    margin-right: 20px;
    .el-link {
      font-weight: normal;
    }
  }
  .layout-header__user {
    cursor: pointer;
    .el-dropdown-link {
      height: 50px;
      display: flex;
      align-items: center;
    }
  }
  .layout-header__brand {
    height: 80px;
    display: flex;
    align-items: center;
    margin-right: auto;
    .page-title {
      margin-left: 10px;
    }
    .nav-menu {
      display: none;
    }
  }
}
// 手机端
body.device-mobile {
  .layout-header {
    padding: 0 $--mobile-padding 0 0;

    .layout-header__brand {
      .nav-menu {
        width: 49px;
        height: 49px;
        display: flex;
        align-items: center;
        justify-content: center;
        box-sizing: border-box;
        background: #eee;
        color: #333;
        .tiger-icon {
          width: 50%;
          height: 50%;
        }
      }
    }
  }
  // 菜单drawer背景色跟菜单背景色一致
  .nav-menu__drawer {
    background-color: #2f3447;
  }
  .nav-menu__brand {
    color: #fff;
    display: flex;
    justify-content: flex-start;
    padding: 10px;
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
</style>
