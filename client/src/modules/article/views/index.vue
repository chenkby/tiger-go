<template>
  <el-skeleton :rows="10" animated :loading="loading">
    <page-header title="活动管理">
      <template #right>
        <el-alert type="success" title="这是说明12312312323"></el-alert>
      </template>
    </page-header>

    <table-header create-title="添加文章" v-model.trim="searchForm.keyword" @search="onSearch">
      <template #inputPrefix>
        <el-select v-model="searchForm.field" placeholder="搜索字段">
          <el-option v-for="item in searchFields" :key="item.value" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
      </template>
      <template #filters>
        <el-form-item>
          <el-select v-model="searchForm.status" placeholder="请选择状态">
          </el-select>
        </el-form-item>
      </template>
    </table-header>

    <el-table :data="tableData" style="width: 100%" size="medium" ref="refTable">
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
    <table-footer v-if="tableData" @select-change="onSelectChange" :pagination="pagination" @current-change="onCurrentChange"></table-footer>

  </el-skeleton>
</template>

<script>
import listMixin from '@/common/mixins/list'
import PageHeader from '@/components/PageHeader.vue'
import TableHeader from '@/components/table/TableHeader.vue'
import TableFooter from '@/components/table/TableFooter.vue'
import { getArticleList, deleteArticle } from '@/modules/article/api'
export default {
  components: { PageHeader, TableHeader, TableFooter },
  data() {
    return {
      searchFields: [
        { label: '用户ID', value: 'user_id' },
        { label: '手机号码', value: 'mobile' }
      ]
    }
  },
  setup() {
    return listMixin(getArticleList, deleteArticle)
  },
  methods: {

  }
}


</script>

<style lang="scss">
.el-table th,
.el-table th.is-leaf {
  background-color: #fafafa;
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
