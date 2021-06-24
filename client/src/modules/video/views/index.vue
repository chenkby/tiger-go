<template>
  <el-skeleton :rows="10" animated :loading="loading">
    <page-header title="活动管理">
      <template #right>
        <el-alert type="success" title="这是说明12312312323"></el-alert>
      </template>
    </page-header>

    <table-header create-title="添加文章" v-model.trim="searchForm.keyword" @search="onSearch">
      <!-- 插到关键词input输入框前面 -->
      <template #prepend></template>
      <!--插到input输入框后面-->
      <template #filters>
        <el-form-item>
          <el-select v-model="searchForm.status" placeholder="请选择状态">
          </el-select>
        </el-form-item>
      </template>
      <!-- 第二行的搜索，如果有很多的筛选条件，可考虑使用频次较少的隐藏起来 -->
      <template #moreFilters>
        <el-form-item>
          <el-input placeholder="这是第二行的搜索"></el-input>
        </el-form-item>
      </template>
    </table-header>

    <el-table :data="tableData" @selection-change="onSelectionChange" size="medium" ref="refTable">
      <template #empty>
        <el-empty :image-size="140">
          <el-button type="primary" icon="el-icon-refresh" @click="refresh">刷新</el-button>
        </el-empty>
      </template>
      <el-table-column type="selection" width="44">
      </el-table-column>
      <el-table-column prop="name" :label="labels.name">
      </el-table-column>
      <download-column prop="video"></download-column>
      <el-table-column prop="size" :label="labels.size">
      </el-table-column>
      <el-table-column prop="play_num" :label="labels.play_num">
      </el-table-column>
      <switch-column></switch-column>
      <!-- <el-table-column prop="free" :label="labels.free">
        <template #default="scope">
          <switch-column :scope="scope"></switch-column>
        </template>
      </el-table-column> -->
      <el-table-column prop="rank" :label="labels.rank">
      </el-table-column>
      <action-column width="150px" @delete="onDelete"></action-column>
    </el-table>

    <table-footer v-if="tableData?.length > 0" :pagination="pagination" @current-change="onCurrentChange" @size-change="onSizeChange" @delete="onDelete"></table-footer>
  </el-skeleton>
  <form-dialog ref="refFormDialog" v-if="dialogMode"></form-dialog>
</template>

<script>
import listMixin from '@/common/mixins/list'
import useList from '@/common/mixins/list/useList'
import { listApi, deleteApi } from './../api'
import { labels } from './../models/video'
import FormDialog from './form.vue'
/**
 * 是否使用对话框模式
 */
const dialogMode = false
export default {
  components: { FormDialog },
  data() {
    return {
      // 搜索字段
      searchFields: [
        { label: '用户ID', value: 'user_id' },
        { label: '手机号码', value: 'mobile' }
      ]
    }
  },
  setup() {
    const searchForm = {
      keyword: '',
      status: ''
    }
    return useList({
      dialogMode: dialogMode,
      listApi: listApi,
      deleteApi: deleteApi,
      primaryKey: 'article_id',
      searchForm: searchForm,
      labels: labels
    })
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
