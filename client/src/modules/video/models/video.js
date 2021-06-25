/**
 * 字段标签
 */
const labels = {
  video_id: '视频ID',
  name: '视频名称',
  category_id: '所属分类',
  image: '封面图',
  video: '视频地址',
  play_num: '播放数',
  status: '状态',
  rank: '排序',
  free: '免费内容',
  size: '视频大小'
}

/**
 * 表单数据验证规则
 */
const rules = {
  name: [{ type: 'string', required: true, max: 50, whitespace: true, fullField: labels.name }],
  category_id: [{ type: 'integer', required: true, min: 1, fullField: labels.category_id }],
  image: [{ type: 'string', required: true, max: 100, whitespace: true, fullField: labels.image }],
  video: [{ type: 'string', required: true, max: 255, whitespace: true, fullField: labels.video }],
  play_num: [{ type: 'integer', required: true, min: 0, fullField: labels.play_num }],
  status: [{ type: 'integer', required: true, min: 0, fullField: labels.status }],
  rank: [{ type: 'integer', required: true, min: 0, fullField: labels.rank }],
  free: [{ type: 'integer', required: true, min: 0, fullField: labels.free }]
}

/**
 * 字段提示
 */
const hints = {
  rank: '数字越大排得越前面'
}

const freeItems = [
  {
    label: '免费',
    value: 1
  },
  {
    label: '收费',
    value: 0
  }
]

export { rules, labels, hints, freeItems }
