import { createStore } from 'vuex'
import table from './modules/table'
import VuexPersistence from 'vuex-persist'
const vuexLocal = new VuexPersistence({
  storage: window.localStorage
})
// 创建一个新的 store 实例
const store = createStore({
  modules: {
    table: table
  },
  plugins: [vuexLocal.plugin]
})
export default store
