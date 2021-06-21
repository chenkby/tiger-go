import parser from 'ua-parser-js'

/**
 * 检测客户端信息
 * @returns 返回检测到的对象
 */
const detect = () => {
  const a = new parser.UAParser()
  return a.getResult()
}

/**
 * 深度合并
 * @param {*} target
 * @param {*} source
 * @returns
 */
const deepMerge = (target, source) => {
  if (source) {
    for (var s in source) {
      if (Object.prototype.hasOwnProperty.call(source, s)) {
        var value = source[s]

        if (typeof value === 'object' && typeof target[s] === 'object') {
          target[s] = Object.assign({}, target[s], value)
        } else {
          target[s] = value
        }
      }
    }
  }
  return target
}

export { detect, deepMerge }
