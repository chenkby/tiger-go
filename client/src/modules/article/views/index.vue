<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>活动管理</el-breadcrumb-item>
    <el-breadcrumb-item>活动列表</el-breadcrumb-item>
    <el-breadcrumb-item>活动详情</el-breadcrumb-item>
  </el-breadcrumb>
  <div class="page-header">
    <div class="title">活动管理</div>
  </div>

  <div class="grid-header">
    <!--query form-->
    <div class="grid-header__left">

      <el-form :inline="true" :model="formInline" class="search-form">
        <div class="el-row">
          <div class="el-col el-col-24">
            <el-form-item>
              <el-input placeholder="输入关键词"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="el-icon-search">查询</el-button>
            </el-form-item>
            <el-form-item>
              <span @click="toggleMoreFilters" class="btn-more-filters">更多筛选项<i class="el-icon-arrow-up"></i></span>
            </el-form-item>
          </div>
        </div>
        <el-collapse-transition>
          <div class="el-row">
            <div class="el-col el-col-24">
              <el-form-item>
                <el-input placeholder="这是第二行的搜索"></el-input>
              </el-form-item>
            </div>
          </div>
        </el-collapse-transition>
      </el-form>

    </div>

    <div class="grid-header__right">
      <el-button icon="el-icon-plus" type="success">添加内容</el-button>
    </div>
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
    <el-table-column fixed="right" label="操作" width="150">
      <el-button type="text">
        编辑
      </el-button>
      <el-divider direction="vertical"></el-divider>
      <el-popover placement="top-end" :width="200" v-model:visible="visible">
        <p>确定要删除这条数据吗？</p>
        <div style="text-align: right; margin: 0">
          <el-button size="mini" type="text" @click="visible = false">取消</el-button>
          <el-button type="primary" size="mini" @click="visible = false">确定</el-button>
        </div>
        <template #reference>
          <el-button type="text" @click="visible = true">删除</el-button>
        </template>
      </el-popover>
      <el-divider direction="vertical"></el-divider>
      <el-button type="text">
        查看
      </el-button>
    </el-table-column>
  </el-table>

  <div class="grid-footer">
    <div class="grid-footer__left">

    </div>
    <div class="grid-footer__right">
      <el-pagination background :page-size="100" layout="total,sizes, prev, pager, next, jumper" :total="1000">
      </el-pagination>
    </div>
  </div>

</template>

<script>
export default {
  data() {
    return {
      formInline: {
        user: '',
        region: '',
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
      }, {
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
      }, {
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
  methods: {

  },
}
</script>

<style lang="scss">
.grid-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  .search-form {
    .btn-more-filters {
      cursor: pointer;
    }
  }
  .grid-header__right {
  }
}
.grid-footer {
  padding: 20px 0 0 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>
