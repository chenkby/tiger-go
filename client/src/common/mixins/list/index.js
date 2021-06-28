/**
 * 列表页的mixin，仅用于components引用
 */
import PageHeader from '@/components/PageHeader.vue'
import Grid from '@/components/table/Grid.vue'
import TableHeader from '@/components/table/TableHeader.vue'
import TableFooter from '@/components/table/TableFooter.vue'
import ActionColumn from '@/components/table/ActionColumn.vue'
import SwitchColumn from '@/components/table/SwitchColumn.vue'
import DownloadColumn from '@/components/table/DownloadColumn.vue'
import GridColumn from '@/components/table/GridColumn.vue'
export default {
  components: { PageHeader, Grid, GridColumn, ActionColumn, TableHeader, TableFooter, SwitchColumn, DownloadColumn }
}
