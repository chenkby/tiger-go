import parser from 'ua-parser-js'

/**
 * 检测客户端信息
 * @returns 返回检测到的对象
 */
const detect = () => {
  const a = new parser.UAParser()
  return a.getResult()
}

export { detect }
