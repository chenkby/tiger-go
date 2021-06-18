<template>
  <div class="tiger-table-header">
    <div class="tiger-table-header__left">
      <slot name="left">
        <el-form :inline="inlineForm" class="search-form" @submit.prevent>
          <el-form-item>
            <el-input class="keyword-input" :placeholder="placeholder" prefix-icon="el-icon-search" clearable v-model="keyword" @input="onKeywordInput">
              <template #prepend v-if="enableInputPrepend">
                <slot name="prepend"></slot>
              </template>
            </el-input>
          </el-form-item>

          <slot name="filters"></slot>
          <!-- 移动端更多筛选项 -->
          <template v-if="$isMobile && showMoreFilters">
            <slot name="moreFilters"></slot>
          </template>
          <!-- end -->

          <el-form-item class="search-form_buttons">
            <span class="space"></span>
            <el-button type="primary" native-type="submit" @click="onSubmitQueryForm" :disabled="submitDisabled">查询</el-button>
            <span class="btn-more-filters">
              <a @click="toggleMoreFilters" v-if="enableMoreFilters">
                更多筛选项<i :class="showMoreFilters ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
              </a>
            </span>
          </el-form-item>
          <!-- PC端更多筛选项 -->
          <el-collapse-transition>
            <div v-if="!$isMobile && showMoreFilters">
              <slot name="moreFilters"></slot>
            </div>
          </el-collapse-transition>
          <!-- end -->
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
import { getCurrentInstance, ref, computed, onMounted, onUnmounted } from 'vue'
import emitter from 'tiny-emitter/instance'
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
  setup(props, { attrs, slots, emit }) {
    const { ctx } = getCurrentInstance()


    // 禁用查询按钮
    const submitDisabled = ref(false)

    const enableInputPrepend = slots.prepend && slots.prepend().length > 0  // input前面是否有内容
    const enableMoreFilters = slots.moreFilters && slots.moreFilters().length > 0  // 根据slot内容判断是否使用了更多筛选
    const showMoreFilters = ref(false)

    const toggleMoreFilters = () => {
      showMoreFilters.value = !showMoreFilters.value
    }

    onMounted(() => {
      emitter.on('loadDataComplate', () => {
        submitDisabled.value = false
      })
    })
    onUnmounted(() => {
      emitter.off('loadDataComplate')
    })


    const keyword = ref('')

    // 输入框改变时
    const onKeywordInput = () => {
      emit('update:modelValue', keyword.value)
    }
    // 提交搜索
    const onSubmitQueryForm = () => {
      submitDisabled.value = true
      emit('search')
    }

    // 根据设备判断搜索表单是否行内表单
    const inlineForm = computed(() => !ctx.$isMobile)
    return {
      inlineForm,
      enableInputPrepend,
      enableMoreFilters,
      showMoreFilters,
      toggleMoreFilters,
      keyword,
      onKeywordInput,
      submitDisabled,
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
    // input前置内容
    .keyword-input {
      .el-input-group__prepend {
        background-color: #fff;
        .el-input--suffix .el-input__inner {
          padding-left: 15px;
        }
        .el-select .el-input {
          width: 120px;
        }
      }
    }
    .btn-more-filters {
      cursor: pointer;
      color: #666;
      display: inline-block;
      margin-left: 15px;
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
      .search-form_buttons {
        .el-form-item__content {
          display: flex;
          align-items: center;
          justify-content: space-between;
          span {
            flex: 1;
          }
        }
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