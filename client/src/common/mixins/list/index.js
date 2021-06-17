/**
 * 列表页的mixin，仅用于component引用
 */
import PageHeader from '@/components/PageHeader.vue'
import TableHeader from '@/components/table/TableHeader.vue'
import TableFooter from '@/components/table/TableFooter.vue'
import ActionColumn from '@/components/table/ActionColumn.vue'
export default {
  components: { PageHeader, ActionColumn, TableHeader, TableFooter }
}
