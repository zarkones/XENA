import Vuex from 'vuex'

import VuexPersist from 'vuex-persistedstate'

const store = () => {
  return new Vuex.Store({
    state: () => ({
      privateKey: '',
      username: '',
      env: {
        XENA_ATILA_HOST: 'http://service.atila.xena.network',
        XENA_PYRAMID_HOST: 'http://service.pyramid.xena.network',
        XENA_RA_HOST: 'http://service.ra.xena.network',
        XENA_XERUM_HOST: 'http://service.xerum.xena.network',
      },
    }),

    plugins: [
      VuexPersist()
    ],

    getters: {
      getUsername: (state) => state.username,
      getPrivateKey: (state) => state.privateKey,
      getAtilaHost: (state) => state.env.XENA_ATILA_HOST,
      getRaHost: (state) => state.env.XENA_RA_HOST,
      getPyramidHost: (state) => state.env.XENA_PYRAMID_HOST,
      getXerumHost: (state) => state.env.XENA_XERUM_HOST,
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

      setAtilaHost: (state, url: string) => {
        state.env.XENA_ATILA_HOST = url
      },

      setRaHost: (state, url: string) => {
        state.env.XENA_RA_HOST = url
      },

      setPyramidHost: (state, url: string) => {
        state.env.XENA_PYRAMID_HOST = url
      },

      setXerumHost: (state, url: string) => {
        state.env.XENA_XERUM_HOST = url
      },
    },

    actions: {
      setPrivateKey: ({ commit }, key: string) => {
        commit('setPrivateKey', key)
      },

      setUsername: ({ commit }, name: string) => {
        commit('setUsername', name)
      },

      setAtilaHost: ({ commit }, url: string) => {
        commit('setAtilaHost', url)
      },

      setRaHost: ({ commit }, url: string) => {
        commit('setRaHost', url)
      },

      setPyramidHost: ({ commit }, url: string) => {
        commit('setPyramidHost', url)
      },

      setXerumHost: ({ commit }, url: string) => {
        commit('setXerumHost', url)
      },
    },
  })
}

export default store