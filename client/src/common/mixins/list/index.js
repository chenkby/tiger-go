/**
 * 列表页的mixin，仅用于components引用
 */
import PageHeader from '@/components/PageHeader.vue'
import TableHeader from '@/components/table/TableHeader.vue'
import TableFooter from '@/components/table/TableFooter.vue'
import ActionColumn from '@/components/table/ActionColumn.vue'
import SwitchColumn from '@/components/table/SwitchColumn.vue'
import DownloadColumn from '@/components/table/DownloadColumn.vue'
export default {
  components: { PageHeader, ActionColumn, TableHeader, TableFooter, SwitchColumn, DownloadColumn }
}
