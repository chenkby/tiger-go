<template>
  <router-view>
    <!-- <el-row>
      <el-col :xs="8" :sm="6" :md="4" :lg="3" :xl="24">header test</el-col>
    </el-row> -->
    <el-container class="layout-container">
      <el-header class="layout__header">
        <div class="app-logo">
          <img class="icon" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAFQklEQVR4nM2aTWgdVRTHf5MvkZSQ6CI0IlJRCcbUjZVCiiC0EN0ZKq3QQhAxDUTthyvXotlIqJu6iCRVkbZB7apSKK4aLXVVa0G0WmyJpUhrldaaR83fxdyxk/fevHfPzJ2nP7i8ZN75uueeO+/eOxMBSKJkOoDzwAMesj8DDwG3ywwoiiIA2sp0kmIMv87j5MZKjGUVEbSkAk4DGwzyXwNPlhQL0NoK2ISt8zj5TSXEUkMrEvB6i/VMlD0FHgG+S/wYETAIfB80IkerpsBu8nUep/dawFgynZRVAfcCl4C7C9i4BdwPXA0SUYpWVMAkxTqP058MEEsmZVXAXcQLmv4Atq4Qrw2WA9j6l7IrYAdhOo+zsyOQrRrKqIAIOAsMBbR5Dhgm/mUIQpkVMErYzuPsjQa2CZSTgL0l2CzNbugpsB44E8pYHR4HvglhqKwpUPbyNbj9kAm4D9hu1Jkzym93foIRMgGvAJ0G+R+Al92nL53OTzBCJWANMGHUmSE+9Zkx6k04f0EIlYAXgV6D/FVg3v09j22t3+v8BSFEAtqJd30WDhBvdHCfB4z6u53fMEgq0rbKxrKk/iob/e66ha1F4k4IUQH7jPIfEm9w0lxx18v0m02BTI4YR02ShjJsDeWwNZI39oSiFWBdnn5OvLGpxzn3fZn+65Mziw9KWjGO2OYmNjcb7a24OP6TCtiL7bzvDHCiicwJbHuJiBBVkCOD90i6YRytnZ62dxrt3nDxtLQCdgHdBvlfgMOesoedvC/dLp5c5ElAFzBl1HkXqHjKVpy8hSniuPJhLJ/xHCXaa/TRK/sUG2/FFIiwL0DeB64bda47PQv7yPsQxpC5LcZRuS1pnXH0k7bO6VvYUnYFWE9kPgMuGHUSLjh9C/lOjDyzNmwcDUnamGPk021jDp/DPrYTLBVgXXB8CZwy6lRzytmxYF8YeWRsraSKcSTGCox8uo0Z/VZcvEErYArbed+PwFGDfCOOOnu+dGJYp/gkoBv7E9oZYMWok8UK9nPDSTxXqj4JGAf6DM6vcee8LxTzzq4vfcRxN6VZAtqAPQbHAO8BN4061fQCj6b+v+nsWtiD7xRv9Q0oo7VLGpV0SNJfkk5VfR/0RuybgEWjw7kcHR+UNC1pqY69wSrZOWM8i1l+fRJQ2iJE8WZnl+JRbsR0lV6wxZhPAhaMjo5n2MkqcR8uO720nePGuBbqxdMsASE3Io1K3Idnq+wF2ZA1S8B+o5OzkqKUvm+J+3CkKrbI+bOwX4YE5D2MyFPiPixL6quKcdxoo+ZQJqHe7+QEtvO+CvGbIReJz/W3Eb8mF4ou4IWqax8Dlw02umn09DqVmS7ln6tlclq10/QNo40lxf1rWAHbgAFDZlvFBmrfPLOuOAeI+7eK6gSEe+AYjiXgbWr3Ann2HPX758rC+liqTP6U9JHimNqUvbbI/XiuXgKOBQq+CCclvSSpR/5L6U+MPo4plYDkPcEh4FtjOYXiEvABcBDbC1MJI8BJo85jURTdeUotaTbI+PnjW+K+7Suj/9mk7x3u8+kcmc/DIvGN6wjwR0C77wALBvm1q/5T/HbGrwVHNYuLkt6U9LCKj3SjjdZPTeKoSDooab0yboJPSPo9SJfDl7hPezUjlmuS3pI0kJavlwAkPeWCz0ueu3iotkbSb6lYzkuaktRdTz4rAUh6Rrajp1aUuG+bVjwIz6lJ5SVkvS7/PHCI7EPFW8CnxDe0Lwh3BF6UduBvH8HkdfmOjO8XgB5gtup6WXfxUHh1Pk1WAiB+Pt9DfLxcZKHyv+YfkNHcqPefgsIAAAAASUVORK5CYII=" />
          <div class="name">COOL-ADMIN</div>
        </div>
        <div class="collapse">
          <i class="el-icon-s-fold"></i>
        </div>

      </el-header>
      <el-container>
        <el-aside width="200px">
          <div class="sidebar-menu">
            <el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" background-color="#2f3447" text-color="#fff" active-text-color="#ffd04b">
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
        </el-aside>
        <el-main>
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>活动管理</el-breadcrumb-item>
            <el-breadcrumb-item>活动列表</el-breadcrumb-item>
            <el-breadcrumb-item>活动详情</el-breadcrumb-item>
          </el-breadcrumb>
          <div class="page-header">
            <div class="title">活动管理</div>
          </div>

          <!--query form-->
          <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item>
              <el-input placeholder="审批人"></el-input>
            </el-form-item>
            <el-form-item>
              <el-select placeholder="活动区域">
                <el-option label="区域一" value="shanghai"></el-option>
                <el-option label="区域二" value="beijing"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary">查询</el-button>
            </el-form-item>
          </el-form>
          <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="date" label="日期" width="180">
            </el-table-column>
            <el-table-column prop="name" label="姓名" width="180">
            </el-table-column>
            <el-table-column prop="address" label="地址">
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>

    </el-container>

  </router-view>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  date() {
    return {
      formInline: {
        user: '',
        region: ''
      },
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄',
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄',
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄',
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄',
      }],
    }
  },
}

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/script-setup-2/active-rfcs/0000-script-setup.md
</script>

<style lang="scss">
body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
}

#app {
  height: 100vh;
}
.layout-container {
  width: 100%;
  height: 100%;
}
.layout__header {
  box-shadow: 0 2px 4px 0 var(--cb-color-shadow, rgba(0, 0, 0, 0.16));
}
#app .el-aside {
  // position: fixed;
  // top: 0;
  // bottom: 0;
  // left: 0;
  // background: #2f3447;
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
    color: #fff;
    font-weight: 700;
    font-size: 24px;
    margin-left: 10px;
  }
}
.sidebar-menu {
  .el-menu {
    border-right: none;
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

.el-breadcrumb {
  font-size: 12px !important;
  margin-bottom: 15px;
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
