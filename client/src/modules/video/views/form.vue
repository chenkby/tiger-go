<template>
  <form-container ref="formContainer" page-title="添加活动">
    <template #left>
      <div class="quick-nav">
        <div class="item title">快速导航</div>
        <div class="item">
          <el-link href="javascript:void 0;" @click="gotoAnchor('panel1')">基本资料</el-link>
        </div>
        <div class="item">
          <el-link href="javascript:void 0;" @click="gotoAnchor('panel2')">内容设置</el-link>
        </div>
        <div class="item">
          <el-link>VIP价格设置</el-link>
        </div>
        <div class="item">
          <el-link>分享设置</el-link>
        </div>
      </div>
    </template>
    <el-form ref="refForm" :model="model" :label-width="labelWidth" :rules="rules" :label-position="labelPosition" @validate="onValidate">

      <field prop="name">
        <el-input v-model="model.name" maxlength="50"></el-input>
      </field>
      <el-row>
        <el-col :span="12">
          <field prop="status">
            <form-radio-group v-model="model.status"></form-radio-group>
          </field>
        </el-col>
        <el-col :span="12">
          <field prop="free">
            <form-radio-group v-model="model.status" :items="boolItems"></form-radio-group>
          </field>
        </el-col>
      </el-row>
      <field prop="rank">
        <el-input v-model.number="model.rank" style="width:120px"></el-input>
      </field>

    </el-form>
    <template #footer>
      <form-buttons @submit="onSubmit" @cancel="onCancel"></form-buttons>
    </template>
  </form-container>
</template>
<script>
import formMixin from '@/common/mixins/form'
import { labels, rules, hints } from './../models/video'
import { createApi, updateApi, infoApi } from './../api'
export default {
  data() {
    return {
      infoApi: infoApi,
      createApi: createApi,
      updateApi: updateApi,
      dialogMode: false,
      labelWidth: '100px',
      rules: rules,
      boolItems: [
        "是", "否"
      ],
      model: {
        status: 1,
        free: 1,
        rank: 0
      }
    }
  },
  provide() {
    return {
      labels: labels
    }
  },
  created() {
    if (this.id) {
      this.loadData({ id: this.id })
    }
  },
  methods: {
    gotoAnchor(name) {
      console.log(`#${name}`)
      document.querySelector(`#${name}`).scrollIntoView(true)
      // var anchor = document.querySelector('#' + name)
      // document.body.scrollTop = anchor.offsetTop
    }
  },
  mixins: [formMixin]
}
</script>
