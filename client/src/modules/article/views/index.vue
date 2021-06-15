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
      <page-header title="活动管理">
        <template #right>
          <el-alert type="success" title="这是说明12312312323"></el-alert>
        </template>
      </page-header>

      <div class="grid-header">
        <!--query form-->
        <!-- <div class="grid-header__left"> -->

        <div class="grid-header__left">
          <!-- <el-button type="primary"> 添加内容</el-button> -->
          <el-form :inline="inlineForm" :model="searchForm" class="search-form" @submit.prevent>

            <el-form-item>
              <el-input placeholder="输入关键词" v-model="searchForm.keyword" prefix-icon="el-icon-search" clearable></el-input>
            </el-form-item>

            <!-- <el-form-item>
                  <span @click="toggleMoreFilters" class="btn-more-filters">
                    更多筛选项<i :class="showMoreFilters ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
                  </span>
                </el-form-item> -->

            <el-collapse-transition>
              <div class="el-row" v-show="showMoreFilters">
                <div class="el-col el-col-24">
                  <el-form-item>
                    <el-input placeholder="这是第二行的搜索"></el-input>
                  </el-form-item>

                </div>
              </div>
            </el-collapse-transition>
            <el-form-item>
              <el-select v-model="searchForm.status" placeholder="请选择状态">
              </el-select>
            </el-form-item>

            <el-form-item class="form-buttons">
              <el-button type="primary" native-type="submit" @click="onSubmitQueryForm">查询</el-button>
            </el-form-item>

          </el-form>
        </div>

        <div class="grid-header__right">
          <el-button type="success" icon="el-icon-plus">添加内容</el-button>
          <!-- <el-button-group style="margin-left:10px;">
            <el-tooltip effect="light" transition="" :append-to-body="true" content="Top Left 提示文字" placement="top-end">
              <el-button icon="el-icon-edit" style="height:40px;"></el-button>
            </el-tooltip>
            <el-tooltip effect="light" transition="" :append-to-body="true" content="刷新页面" placement="top-start">
              <el-button icon="el-icon-share" style="height:40px;"></el-button>
            </el-tooltip>
            <el-tooltip effect="light" transition="" :append-to-body="true" content="导出数据" placement="top-end">
              <el-button icon="el-icon-delete" style="height:40px;"></el-button>
            </el-tooltip>
          </el-button-group> -->
        </div>

        <!-- </div> -->

        <!-- <el-col :xl="12" class="grid-header__right">
            <el-button icon="el-icon-plus" type="success">添加内容</el-button>
          </el-col> -->

      </div>

      <el-table :data="tableData" style="width: 100%" size="medium">
        <template #empty>
          <el-empty :image-size="100"></el-empty>
        </template>
        <el-table-column type="selection" width="44">
        </el-table-column>
        <el-table-column prop="title" label="日期" width="100">
        </el-table-column>
        <el-table-column prop="name" label="姓名" width="100">
        </el-table-column>
        <el-table-column prop="address" label="地址">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="150" class-name="action-column">
          <template #default="scope">
            <router-link :to="{path: '/article/create'}">编辑</router-link>
            <el-divider direction="vertical"></el-divider>
            <el-popconfirm placement="top-start" @confirm="onConfirmDelete(scope.row)" confirmButtonText='删除' confirmButtonType="danger" icon="el-icon-info" iconColor="red" title="确定要删除这条数据吗？">
              <template #reference>
                <span>删除</span>
              </template>
            </el-popconfirm>
            <el-divider direction="vertical"></el-divider>

            <router-link :to="{path: '/article/view', query:{id:scope.row.article_id}}">查看</router-link>
          </template>
        </el-table-column>
      </el-table>
      <table-footer v-if="tableData" :pagination="pagination" @current-change="onCurrentChange"></table-footer>

    </template>
  </el-skeleton>
</template>

<script setup>
import listMixin from '@/common/mixins/list'
import PageHeader from '@/components/PageHeader.vue'
import TableFooter from '@/components/table/TableFooter.vue'
import { getArticleList, deleteArticle } from '@/modules/article/api'
let { loading, searchForm, inlineForm, tableData, pagination, onCurrentChange, onSubmitQueryForm } = listMixin(getArticleList, deleteArticle)
</script>

<style lang="scss">
.el-table th,
.el-table th.is-leaf {
  background-color: #fafafa;
}
.grid-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  //margin-bottom: 20px;
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
  .grid-header__left {
    display: flex;
    align-items: center;

    .el-form-item {
      //margin-bottom: 0;
    }
  }
  .grid-header__right {
    justify-content: space-between;
    display: flex;
  }
}
// 手机端
body.device-mobile {
  .grid-header {
    display: block;
    .search-form {
      width: 100%;
      .el-select {
        width: 100%;
      }
      .form-buttons {
        text-align: center;
      }
    }
    .grid-header__left {
      margin-bottom: 20px;
      .el-form-item {
        margin-bottom: 15px;
      }
    }
    .grid-header__right {
      text-align: left;
      margin-bottom: 20px;
    }
  }

  .grid-footer {
    display: block;
    .grid-footer__left {
    }
    .grid-footer__tools {
    }
    .grid-footer__right {
      text-align: center;
    }
    .el-pagination {
      border-top: 1px solid #efefef;
      padding-top: 20px;
      margin-top: 20px;
    }
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
