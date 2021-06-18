<template>
  <el-skeleton :rows="10" animated :loading="loading">
    <page-header title="活动管理">
      <template #right>
        <el-alert type="success" title="这是说明12312312323"></el-alert>
      </template>
    </page-header>

    <table-header create-title="添加文章" v-model.trim="searchForm.keyword" @search="onSearch">
      <template #prepend>

      </template>
      <template #filters>
        <el-form-item>
          <el-select v-model="searchForm.status" placeholder="请选择状态">
          </el-select>
        </el-form-item>
      </template>

      <template #moreFilters>
        <el-form-item>
          <el-input placeholder="这是第二行的搜索"></el-input>
        </el-form-item>

      </template>
    </table-header>

    <el-table :data="tableData" style="width: 100%" size="medium" ref="refTable">
      <template #empty>
        <el-empty :image-size="140">
          <el-button type="primary" icon="el-icon-refresh">刷新</el-button>
        </el-empty>
      </template>
      <el-table-column type="selection" width="44">
      </el-table-column>
      <el-table-column prop="title" label="日期">
      </el-table-column>
      <el-table-column prop="name" label="姓名" width="100">
      </el-table-column>
      <el-table-column prop="address" label="地址">
      </el-table-column>
      <action-column width="150px" @update="onUpdate" @delete="onDelete"></action-column>
    </el-table>
    <table-footer v-if="tableData" @select-change="onSelectChange" :pagination="pagination" @current-change="onCurrentChange"></table-footer>

  </el-skeleton>
</template>

<script>
import listMixin from '@/common/mixins/list'
import useList from '@/common/mixins/list/useList'
import { getArticleList, deleteArticle } from '@/modules/article/api'
import { ref, onMounted, provide } from 'vue'
export default {
  data() {
    return {
      searchFields: [
        { label: '用户ID', value: 'user_id' },
        { label: '手机号码', value: 'mobile' }
      ]
    }
  },
  setup() {
    const table = ref(null)
    const abc = ref(null)
    provide('primaryKey', 'article_id')
    onMounted(() => {
      // DOM元素将在初始渲染后分配给ref
      console.log('tab', table.value, abc) // <div>这是根元素</div>
    })
    return useList(getArticleList, deleteArticle)
  },

  mixins: [listMixin]
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
.el-table__empty-text {
  .el-empty__description p {
    line-height: 1;
  }
}
</style>
