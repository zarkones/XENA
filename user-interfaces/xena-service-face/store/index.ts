import Vuex from 'vuex'

import VuexPersist from 'vuex-persistedstate'

const store = () => {
  return new Vuex.Store({
    state: () => ({
      privateKey: '',
      username: '',
    }),

    plugins: [
      VuexPersist()
    ],

    getters: {
      getUsername: (state) => state.username,
      getPrivateKey: (state) => state.privateKey,
    },

    mutations: {
      setPrivateKey: (state, key: string) => {
        state.privateKey = key.replaceAll(' ', '\n')
          .replace(`-----BEGIN\nRSA\nPRIVATE\nKEY-----`, '-----BEGIN RSA PRIVATE KEY-----')
          .replace(`-----END\nRSA\nPRIVATE\nKEY-----`, '-----END RSA PRIVATE KEY-----')
      },

      setUsername: (state, name: string) => {
        state.username = name
      },
    },

    actions: {
      setPrivateKey: ({ commit }, key: string) => {
        commit('setPrivateKey', key)
      },

      setUsername: ({ commit }, name: string) => {
        commit('setUsername', name)
      },
    },
  })
}

export default store