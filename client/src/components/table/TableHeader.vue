<template>
  <div class="tiger-table-header">
    <div class="tiger-table-header__left">
      <slot name="left">
        <el-form :inline="inlineForm" class="search-form" @submit.prevent>
          <el-form-item>
            <el-input :placeholder="placeholder" prefix-icon="el-icon-search" clearable v-model="keyword" @input="onKeywordInput">
              <slot name="inputPrefix">
                <div>
                  <template #prepend>aaa</template>
                </div>
              </slot>

            </el-input>
          </el-form-item>

          <!-- <el-form-item>
                  <span @click="toggleMoreFilters" class="btn-more-filters">
                    更多筛选项<i :class="showMoreFilters ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
                  </span>
                </el-form-item> -->

          <!-- <el-collapse-transition>
            <div class="el-row" v-show="showMoreFilters">
              <div class="el-col el-col-24">
                <el-form-item>
                  <el-input placeholder="这是第二行的搜索"></el-input>
                </el-form-item>

              </div>
            </div>
          </el-collapse-transition> -->
          <slot name="filters"></slot>
          <el-form-item class="form-buttons">
            <el-button type="primary" native-type="submit" @click="onSubmitQueryForm">查询</el-button>
          </el-form-item>

        </el-form>
      </slot>
    </div>

    <div class="tiger-table-header__right">
      <slot name="right">
        <el-button type="success" icon="el-icon-plus">{{createTitle}}</el-button>
        <slot name="tool">
          <!-- <el-button-group style="margin-left:10px;">
            <el-tooltip effect="light" transition="" :append-to-body="true" content="Top Left 提示文字" placement="top-end">
              <el-button icon="el-icon-edit"></el-button>
            </el-tooltip>
            <el-tooltip effect="light" transition="" :append-to-body="true" content="刷新页面" placement="top-start">
              <el-button icon="el-icon-share"></el-button>
            </el-tooltip>
            <el-tooltip effect="light" transition="" :append-to-body="true" content="导出数据" placement="top-end">
              <el-button icon="el-icon-delete"></el-button>
            </el-tooltip>
          </el-button-group> -->
        </slot>
      </slot>
    </div>
  </div>
</template>

<script>
import { getCurrentInstance, ref, computed } from 'vue'
export default {
  props: {
    modelValue: String,
    placeholder: {
      type: String,
      default: '输入关键词'
    },
    createTitle: {
      type: String,
      default: '添加内容'
    }
  },
  emits: [
    'update:modelValue',
    'search'
  ],
  setup(props, { emit }) {
    const { ctx } = getCurrentInstance()

    const keyword = ref('')

    // 输入框改变时
    const onKeywordInput = () => {
      emit('update:modelValue', keyword.value)
    }
    // 提交搜索
    const onSubmitQueryForm = () => {
      emit('search')
    }

    // 根据设备判断搜索表单是否行内表单
    const inlineForm = computed(() => !ctx.$isMobile)
    return {
      inlineForm,
      keyword,
      onKeywordInput,
      onSubmitQueryForm
    }
  }
}
</script>


<style lang="scss">
.tiger-table-header {
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
  .tiger-table-header__left {
    display: flex;
    align-items: center;

    .el-form-item {
      //margin-bottom: 0;
    }
  }
  .tiger-table-header__right {
    justify-content: space-between;
    display: flex;
  }
}

// 手机端
body.device-mobile {
  .tiger-table-header {
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
    .tiger-table-header__left {
      margin-bottom: 20px;
      .el-form-item {
        margin-bottom: 15px;
      }
    }
    .tiger-table-header__right {
      text-align: left;
      margin-bottom: 20px;
    }
  }
}
</style>