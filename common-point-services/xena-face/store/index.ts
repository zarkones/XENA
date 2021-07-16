import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const vuex = () => new Vuex.Store({
  state: () => ({
    // Used for signing messages.
    privateKey: '',
  }),
  mutations: {
    setPrivateKey: (state, key) => {
      state.privateKey = key
    }
  },
  actions: {
    setPrivateKey: (state, key) => {
      state.commit('setPrivateKey', key)
    }
  },
  modules: {},
  getters: {
    getPrivateKey: state => state.privateKey
  },
})

export default vuex