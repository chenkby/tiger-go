const table = {
  state: () => {
    return {
      pageSize: {}
    }
  },
  mutations: {
    setPageSize(state, payload) {
      const { path, pageSize } = payload
      state.pageSize[path] = pageSize
    }
  },
  actions: {
    setPageSize({ commit }, payload) {
      commit('setPageSize', payload)
    }
  }
}
export default table
