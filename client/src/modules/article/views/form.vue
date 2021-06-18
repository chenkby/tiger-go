<template>
  <form-container>
    <el-form ref="form" :model="form" label-width="100px" :label-position="labelPosition">
      <panel title="基本资料" name="panel1">
        <el-form-item label="活动名称：">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="活动区域">
          <el-select v-model="form.region" placeholder="请选择活动区域">
            <el-option label="区域一" value="shanghai"></el-option>
            <el-option label="区域二" value="beijing"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="活动时间">
          <el-col :span="11">
            <el-date-picker type="date" placeholder="选择日期" v-model="form.date1" style="width: 100%;"></el-date-picker>
          </el-col>
          <el-col class="line" :span="2">-</el-col>
          <el-col :span="11">
            <el-time-picker placeholder="选择时间" v-model="form.date2" style="width: 100%;"></el-time-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="即时配送">
          <el-switch v-model="form.delivery"></el-switch>
        </el-form-item>
        <el-form-item label="活动性质">
          <el-checkbox-group v-model="form.type">
            <el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>
            <el-checkbox label="地推活动" name="type"></el-checkbox>
            <el-checkbox label="线下主题活动" name="type"></el-checkbox>
            <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="特殊资源">
          <el-radio-group v-model="form.resource">
            <el-radio label="线上品牌商赞助"></el-radio>
            <el-radio label="线下场地免费"></el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="活动形式">
          <el-input type="textarea" v-model="form.desc"></el-input>
        </el-form-item>
      </panel>

    </el-form>
    <template #footer>
      <form-buttons :label-width="labelWidth" @submit="onSubmit" @cancel="onCancel"></form-buttons>
    </template>
  </form-container>
  <!-- <div class="form-container">
    <el-row>
      <el-col :lg="3">
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
      </el-col>
      <el-col :lg="15">
        <el-form ref="form" :model="form" label-width="100px" :label-position="labelPosition">
          <panel title="基本资料" name="panel1">
            <el-form-item label="活动名称：">
              <el-input v-model="form.name"></el-input>
            </el-form-item>
            <el-form-item label="活动区域">
              <el-select v-model="form.region" placeholder="请选择活动区域">
                <el-option label="区域一" value="shanghai"></el-option>
                <el-option label="区域二" value="beijing"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="活动时间">
              <el-col :span="11">
                <el-date-picker type="date" placeholder="选择日期" v-model="form.date1" style="width: 100%;"></el-date-picker>
              </el-col>
              <el-col class="line" :span="2">-</el-col>
              <el-col :span="11">
                <el-time-picker placeholder="选择时间" v-model="form.date2" style="width: 100%;"></el-time-picker>
              </el-col>
            </el-form-item>
            <el-form-item label="即时配送">
              <el-switch v-model="form.delivery"></el-switch>
            </el-form-item>
            <el-form-item label="活动性质">
              <el-checkbox-group v-model="form.type">
                <el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>
                <el-checkbox label="地推活动" name="type"></el-checkbox>
                <el-checkbox label="线下主题活动" name="type"></el-checkbox>
                <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="特殊资源">
              <el-radio-group v-model="form.resource">
                <el-radio label="线上品牌商赞助"></el-radio>
                <el-radio label="线下场地免费"></el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="活动形式">
              <el-input type="textarea" v-model="form.desc"></el-input>
            </el-form-item>
          </panel>



        </el-form>
      </el-col>
    </el-row>

  </div> -->
</template>
<script>
import useForm from '@/common/mixins/form/useForm'
import formMixin from '@/common/mixins/form'

export default {
  data() {
    return {
      labelWidth: '100px',
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
  setup() {
    return useForm()
  },
  created() {
    console.log(this.$route.query)
  },
  provide: {
    dialogMode: false
  },
  methods: {
    onSubmit() {
      console.log('submit!')
    },
    gotoAnchor(name) {
      console.log(`#${name}`)
      document.querySelector(`#${name}`).scrollIntoView(true)
      // var anchor = document.querySelector('#' + name)
      // document.body.scrollTop = anchor.offsetTop
    }
  },
  computed: {
    /**
     * 表单label位置
     */
    labelPosition() {
      return this.$device === 'mobile' ? 'top' : 'right'
    }
  },
  mixins: [formMixin]
}
</script>

<style lang="scss">
body.device-mobile {
  // 手机模式下，表单label下边距
  .el-form--label-top .el-form-item__label {
    padding: 0;
    line-height: 30px;
  }
}

.form-container {
  margin-top: 20px;
  position: relative;

  .quick-nav {
    border-right: 1px solid #efefef;
    text-align: right;
    padding-right: 20px;
    width: 120px;
    position: fixed;

    .item {
      padding: 5px 0;
      font-size: 12px;
      .el-link.el-link--default {
        color: #333;
        &:hover {
          color: $--color-primary;
        }
      }
    }
    .title {
      font-size: 16px;
      margin-bottom: 5px;
      color: #777;
    }
    .el-link {
      font-weight: normal;
      font-size: 80%;
    }
  }
}
</style>
