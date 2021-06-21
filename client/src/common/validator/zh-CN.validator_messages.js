const messages = () => {
  return {
    default: 'Validation error on field %s',
    required: '%s不能为空',
    enum: '%s must be one of %s',
    whitespace: '%s不能为空',
    date: {
      format: '%s date %s is invalid for format %s',
      parse: '%s date could not be parsed, %s is invalid ',
      invalid: '%s date %s is invalid'
    },
    types: {
      string: (label) => {
        return label + '不是一个字符串'
      },
      method: '%s is not a %s (function)',
      array: (label) => {
        return label + '不是一个数组'
      },
      object: '%s is not an %s',
      number: (label) => {
        return label + '必须是一个数字'
      },
      date: (label) => {
        return label + '必须是日期格式'
      },
      boolean: '%s is not a %s',
      integer: (label) => {
        return label + '必须是整数'
      },
      float: '%s is not a %s',
      regexp: '%s is not a valid %s',
      email: '%s is not a valid %s',
      url: (label) => {
        return label + '不是有效的URL，必须包含http(s)://'
      },
      hex: '%s is not a valid %s'
    },
    string: {
      len: '%s must be exactly %s characters',
      min: '%s不能少于%s个字符',
      max: '%s不能超过%s个字符',
      range: '%s必须在%s至%s个字符之间'
    },
    number: {
      len: '%s must equal %s',
      min: '%s不能小于%s',
      max: '%s cannot be greater than %s',
      range: '%s只能在%s 至 %s之间'
    },
    array: {
      len: '%s must be exactly %s in length',
      min: '%s cannot be less than %s in length',
      max: '%s cannot be greater than %s in length',
      range: '%s must be between %s and %s in length'
    },
    pattern: {
      mismatch: '%s value %s does not match pattern %s'
    },
    clone: function clone() {
      let cloned = JSON.parse(JSON.stringify(this))
      cloned.clone = this.clone
      return cloned
    }
  }
}

export default messages()
