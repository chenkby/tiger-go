<template>
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
</template>

<script>
export default {

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
  .app-logo {
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: auto;
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
