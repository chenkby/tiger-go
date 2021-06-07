<template>
  <el-skeleton :rows="10" animated :loading="loading">
    <!-- <template #template>
      <el-skeleton-item variant="cert" style="width: 240px; height: 20px;" />
      <div>
        <el-skeleton-item variant="h1" style="width: 100px; height: 40px;" />
      </div>
      <div style="display:flex;align-items:center;justify-content: space-between;">
        <el-skeleton-item variant="rect" style="width:300px;height: 40px;" />
        <el-skeleton-item variant="button" style="width: 100px;height: 40px;" />
      </div>
      <el-skeleton-item variant="p" style="width: 50%" />
      <div style="display: flex; align-items: center; justify-items: space-between;">
        <el-skeleton-item variant="text" style="margin-right: 16px;" />
        <el-skeleton-item variant="text" style="width: 30%;" />
      </div>

    </template> -->

    <template #default>
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>活动管理</el-breadcrumb-item>
        <el-breadcrumb-item>活动列表</el-breadcrumb-item>
        <el-breadcrumb-item>活动详情</el-breadcrumb-item>
      </el-breadcrumb>
      <div class="page-header">
        <div class="title">活动管理2</div>
      </div>

      <div class="grid-header">
        <!--query form-->
        <!-- <div class="grid-header__left"> -->
        <el-row>
          <el-col :xl="12">
            <el-form :inline="inlineForm" :model="formInline" class="search-form">
              <el-row>
                <el-col :xl="24">
                  <el-form-item>
                    <el-input placeholder="输入关键词" prefix-icon="el-icon-search" clearable></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-select placeholder="请选择状态">
                    </el-select>
                  </el-form-item>
                  <el-form-item>
                    <el-select placeholder="请选择状态">
                    </el-select>
                  </el-form-item>
                  <el-form-item class="form-buttons">
                    <el-button type="primary">查询</el-button>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-form-item>
                <span @click="toggleMoreFilters" class="btn-more-filters">
                  更多筛选项<i :class="showMoreFilters ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
                </span>
              </el-form-item>

              <el-collapse-transition>
                <div class="el-row" v-show="showMoreFilters">
                  <div class="el-col el-col-24">
                    <el-form-item>
                      <el-input placeholder="这是第二行的搜索"></el-input>
                    </el-form-item>

                  </div>
                </div>
              </el-collapse-transition>
            </el-form>
          </el-col>

          <!-- </div> -->

          <!-- <div class="grid-header__right"> -->
          <el-col :xl="12" class="grid-header__right">
            <el-button icon="el-icon-plus" type="success">添加内容</el-button>
          </el-col>
          <!-- </div> -->
        </el-row>
      </div>

      <el-table :data="tableData" style="width: 100%" size="medium">
        <el-table-column type="selection" width="55">
        </el-table-column>
        <el-table-column prop="date" label="日期" width="180">
        </el-table-column>
        <el-table-column prop="name" label="姓名" width="180">
        </el-table-column>
        <el-table-column prop="address" label="地址">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="150" class-name="action-column">
          <router-link :to="{path: '/article/create'}">编辑</router-link>
          <el-divider direction="vertical"></el-divider>
          <el-popconfirm placement="top-start" confirmButtonText='删除' confirmButtonType="danger" icon="el-icon-info" iconColor="red" title="确定要删除这条数据吗？">
            <template #reference>
              <span>删除</span>
            </template>
          </el-popconfirm>
          <el-divider direction="vertical"></el-divider>
          <router-link :to="{path: '/article/view'}">查看</router-link>
        </el-table-column>
      </el-table>

      <div class="grid-footer">
        <div class="grid-footer__left">

        </div>
        <div class="grid-footer__right">
          <el-pagination background :page-size="100" :layout="paginationLayout" :total="1000">
          </el-pagination>
        </div>
      </div>
    </template>
  </el-skeleton>
</template>

<script>
export default {
  data() {
    return {
      showMoreFilters: false,
      loading: true,
      formInline: {
        user: '',
        region: ''
      },
      options: [{
        value: '选项1',
        label: '黄金糕'
      }, {
        value: '选项2',
        label: '双皮奶'
      }, {
        value: '选项3',
        label: '蚵仔煎'
      }, {
        value: '选项4',
        label: '龙须面'
      }, {
        value: '选项5',
        label: '北京烤鸭'
      }],
      value: '',

      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }, {
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }]
    }
  },
  methods: {
    toggleMoreFilters() {
      this.showMoreFilters = !this.showMoreFilters
    }
  },
  created() {
    setTimeout(() => {
      this.loading = false
    }, 300)
  },
  computed: {
    /**
     * 分页器布局
     */
    paginationLayout() {
      if (this.$device === 'mobile') {
        return 'total, prev, next, jumper'
      }
      return 'total,sizes, prev, pager, next, jumper'
    },
    /**
     * 搜索表单是否行内表单
     */
    inlineForm() {
      return this.$device !== 'mobile'
    }
  }
}
</script>

<style lang="scss">
.grid-header {
  // display: flex;
  // align-items: flex-start;
  // justify-content: space-between;
  .search-form {
    .btn-more-filters {
      cursor: pointer;
      color: #666;
      &:hover {
        color: #222;
      }
      i {
        margin-left: 3px;
      }
    }
  }
  .grid-header__right {
    text-align: right;
  }
}
body.device-mobile {
  .grid-header {
    .search-form {
      .el-select {
        width: 100%;
      }
      .form-buttons {
        text-align: center;
      }
    }
    .grid-header__right {
      text-align: left;
      margin-bottom: 20px;
    }
  }
}
.grid-footer {
  padding: 20px 0 0 0;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .el-pagination {
    font-weight: normal;
  }
}
.el-table {
  color: #111;
}
// 操作列
.action-column {
  text-align: right !important;
  .el-link,
  .el-button {
    font-weight: normal;
  }
  .el-link {
    line-height: 1;
  }
  span,
  a {
    color: $--color-primary;
    text-decoration: none;
    cursor: pointer;
    &:hover {
      text-decoration: underline;
    }
  }
}
</style>
