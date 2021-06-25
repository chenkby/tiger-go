import FormContainer from '@/components/form/FormContainer.vue'
import PageHeader from '@/components/PageHeader.vue'
import Panel from '@/components/Panel.vue'
import Field from '@/components/form/Field.vue'
import FormRadioGroup from '@/components/form/FormRadioGroup.vue'
import FormButtons from '@/components/form/FormButtons.vue'
import emitter from 'tiny-emitter/instance'
import { infoApi } from '@/modules/video/api/index'

export default {
  components: { FormContainer, PageHeader, Panel, FormButtons, Field, FormRadioGroup },
  data() {
    return {
      // ID
      id: 0,
      // 加载数据中
      loading: false,
      // model
      model: {},
      // 服务器返回的response
      response: {},
      // 是否对话框模式
      dialogMode: true,
      // 接口URL get参数
      queryParams: {},
      labels: {}
    }
  },
  provide() {
    return {
      dialogMode: this.dialogMode
    }
  },
  created() {
    this.id = this.$route.query.id || 0
  },
  computed: {
    /**
     * 表单label宽度
     * @returns {string}
     */
    labelWidth() {
      return this.dialogMode ? '120px' : '200px'
    },
    /**
     * 返回action
     * @returns {string}
     */
    action() {
      return this.id ? 'update' : 'create'
    },
    /**
     * 提交按钮文本
     * @returns {string}
     */
    submitTitle() {
      return this.id ? '保存' : '修改'
    },
    /**
     * 标签显示位置
     * @returns
     */
    labelPosition() {
      return this.$isMobile ? 'top' : 'right'
    }
  },
  methods: {
    /**
     * 验证时通知这个字段
     * @param {*} prop
     * @param {*} isPass
     * @param {*} errMessage
     */
    onValidate(prop, isPass, errMessage) {
      emitter.emit(`onValidated:${prop}`, errMessage)
    },
    /**
     * 提交表单
     * @param message 成功后提示消息
     */
    async onSubmit(message) {
      return this.$refs.form
        .validate()
        .then(async (s) => {
          let res = {}
          let queryParams = {}
          if (!this.id) {
            res = await this.createApi.call(this, this.model, queryParams)
          } else {
            res = await this.updateApi.call(this, this.model, queryParams)
          }

          if (res.err_code === 0) {
            this.$message({
              message: message || (this.id ? '修改成功' : '创建成功'),
              type: 'success'
            })
            if (this.dialogMode) {
              this.$refs.container.close()
              this.$emit('success', res)
              this.$refs.form.resetFields()
            } else {
              this.$router.go(-1)
            }
          }
          return res
        })
        .catch((err) => {})
    },

    /**
     * 加载model,编辑时使用返回model,添加时返回初始值
     * @param id
     */
    async loadData(queryParams) {
      this.loading = true
      const res = await infoApi(queryParams)
      this.loading = false
      if (res.err_code === 0) {
        this.formData = { ...this.formData, ...res.data }
      }

      //this.closeLoading()
    },
    /**
     * 返回queryParams
     * @returns {queryParams|{}}
     */
    getQueryParams() {
      let queryParams = this.queryParams || {}
      if (this.id) {
        queryParams = { ...queryParams, ...{ id: this.id } }
      }
      return queryParams
    },
    /**
     * 关闭正在加载中的动画
     */
    closeLoading() {
      this.$nextTick(() => {
        // 以服务的方式调用的 Loading 需要异步关闭
        this.loading = false
      })
    },
    /**
     * 取消，返回上一页
     */
    onCancel() {
      if (this.dialogMode) {
        this.$refs.container.close()
      } else {
        this.$router.go(-1)
      }
    },
    /**
     * 打开表单对话框
     * @param id 编辑的主键值
     * @param params
     */
    openDialog(id, params) {
      params.value = params || {}
      this.id = id || ''
      if (params.model) {
        this.model = params.model
      }
      this.$refs.formContainer.open()
    },
    /**
     * 打开对话框初始化数据
     */
    onOpen() {
      if (this.initData) {
        this.initData.call(this)
      }
    },
    /**
     * 关闭对话框初始化数据,清除验证状态
     */
    onClosed() {
      this.model = this.$options.data.call(this).model || {}
      this.$refs.form.clearValidate()
      this.$refs.form.resetFields()
    },
    /**
     * 返回路由信息
     */
    getBreadcrumb() {
      this.breadcrumbs = this.$route.matched
    }
  }
}
