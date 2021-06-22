const labels = {
  activity_id: '活动编号',
  name: '活动名称',
  image: '封面图',
  in_app: '显示的app',
  show_rule: '显示规则',
  link_type: '链接数据',
  status: '状态',
  rank: '排序',
  created_at: '创建时间',
  updated_at: '修改时间',
  link_url: 'H5链接',
  link_page: '小程序打开页面',
  mp_url: '公众号链接',
  type: '活动类型',
  channels: '显示的渠道',
  enable_mina: '是否跳转小程序'
}

const rules = {
  name: [
    {
      type: 'string',
      required: true,
      max: 50,
      whitespace: true,
      fullField: labels.name
    }
  ],
  image: [
    {
      type: 'string',
      required: true,
      max: 255,
      whitespace: true,
      fullField: labels.image
    }
  ],
  in_app: [
    {
      type: 'array',
      required: true,
      whitespace: true,
      fullField: labels.in_app
    }
  ],
  show_rule: [
    {
      type: 'string',
      required: true,
      max: 30,
      whitespace: true,
      fullField: labels.show_rule
    }
  ],
  link_page: [
    {
      type: 'string',
      required: true,
      max: 50,
      whitespace: true,
      fullField: labels.link_page
    }
  ],
  type: [
    {
      type: 'string',
      required: true,
      max: 50,
      whitespace: true,
      fullField: labels.type
    }
  ],
  status: [
    { type: 'integer', required: true, min: 0, fullField: labels.status }
  ],
  enable_mina: [
    { type: 'integer', required: true, min: 0, fullField: labels.enable_mina }
  ],
  rank: [{ type: 'integer', required: true, min: 0, fullField: labels.rank }],
  link_url: [
    {
      type: 'string',
      required: true,
      max: 255,
      whitespace: true,
      fullField: labels.link_url
    }
  ],
  desc: [
    {
      type: 'string',
      required: true,
      max: 255,
      whitespace: true,
      fullField: labels.mp_url
    }
  ]
}

const linkType = [
  { label: '小程序', value: 'mina' },
  { label: '邀请活动', value: 'invite' }
]

const hints = {
  rank: '数字越大排得越前面'
}

export { rules, labels, hints, linkType }
