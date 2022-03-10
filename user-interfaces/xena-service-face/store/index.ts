import Vuex from 'vuex'

import VuexPersist from 'vuex-persistedstate'

const store = () => {
  return new Vuex.Store({
    state: () => ({
      privateKey: '',
      username: '',
      env: {
        XENA_ATILA_HOST: 'http://service.atila.xena.network/v1',
        XENA_ATILA_TOKEN: '',
        
        XENA_DOMENA_HOST: 'http://service.domena.xena.network/v1',
        XENA_DOMENA_TOKEN: '',

        XENA_PYRAMID_HOST: 'http://service.pyramid.xena.network/v1',
        XENA_PYRAMID_TOKEN: '',

        XENA_RA_HOST: 'http://service.ra.xena.network/v1',
        XENA_RA_TOKEN: '',

        XENA_XERUM_HOST: 'http://service.xerum.xena.network/v1',
        XENA_XERUM_TOKEN: '',

        XENA_SENSI_HOST: 'http://service.sensi.xena.network/v1',
        XENA_SENSI_TOKEN: '',
      },

      directBotConnection: '',
    }),

    plugins: [
      VuexPersist()
    ],

    getters: {
      getBotHost: (state) => state.directBotConnection,

      getUsername: (state) => state.username,
      getPrivateKey: (state) => state.privateKey,
      
      getAtilaHost: (state) => state.env.XENA_ATILA_HOST,
      getAtilaToken: (state) => state.env.XENA_ATILA_TOKEN,

      getRaHost: (state) => state.env.XENA_RA_HOST,
      getRaToken: (state) => state.env.XENA_RA_TOKEN,

      getPyramidHost: (state) => state.env.XENA_PYRAMID_HOST,
      getPyramidToken: (state) => state.env.XENA_PYRAMID_TOKEN,

      getXerumHost: (state) => state.env.XENA_XERUM_HOST,
      getXerumToken: (state) => state.env.XENA_XERUM_TOKEN,

      getDomenaHost: (state) => state.env.XENA_DOMENA_HOST,
      getDomenaToken: (state) => state.env.XENA_DOMENA_TOKEN,

      getSensiHost: (state) => state.env.XENA_SENSI_HOST,
      getSensiToken: (state) => state.env.XENA_SENSI_TOKEN,
    },

    mutations: {
      setBotHost: (state, host: string) => state.directBotConnection = host,

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
      setAtilaToken: (state, token: string) => {
        state.env.XENA_ATILA_TOKEN = token
      },

      setRaHost: (state, url: string) => {
        state.env.XENA_RA_HOST = url
      },
      setRaToken: (state, token: string) => {
        state.env.XENA_RA_TOKEN = token
      },

      setPyramidHost: (state, url: string) => {
        state.env.XENA_PYRAMID_HOST = url
      },
      setPyramidToken: (state, token: string) => {
        state.env.XENA_PYRAMID_TOKEN = token
      },

      setXerumHost: (state, url: string) => {
        state.env.XENA_XERUM_HOST = url
      },
      setXerumToken: (state, token: string) => {
        state.env.XENA_XERUM_TOKEN = token
      },

      setDomenaHost: (state, url: string) => {
        state.env.XENA_DOMENA_HOST = url
      },
      setDomenaToken: (state, token: string) => {
        state.env.XENA_DOMENA_TOKEN = token
      },

      setSensiHost: (state, url: string) => {
        state.env.XENA_SENSI_HOST = url
      },
      setSensiToken: (state, token: string) => {
        state.env.XENA_SENSI_TOKEN = token
      },
    },

    actions: {
      setBotHost: ({ commit }, host: string) => commit('setBotHost', host),

      setPrivateKey: ({ commit }, key: string) => {
        commit('setPrivateKey', key)
      },
      setUsername: ({ commit }, name: string) => {
        commit('setUsername', name)
      },

      setAtilaHost: ({ commit }, url: string) => {
        commit('setAtilaHost', url)
      },
      setAtilaToken: ({ commit }, token: string) => {
        commit('setAtilaToken', token)
      },

      setRaHost: ({ commit }, url: string) => {
        commit('setRaHost', url)
      },
      setRaToken: ({ commit }, token: string) => {
        commit('setRaToken', token)
      },

      setPyramidHost: ({ commit }, url: string) => {
        commit('setPyramidHost', url)
      },
      setPyramidToken: ({ commit }, token: string) => {
        commit('setPyramidToken', token)
      },

      setXerumHost: ({ commit }, url: string) => {
        commit('setXerumHost', url)
      },
      setXerumToken: ({ commit }, token: string) => {
        commit('setXerumToken', token)
      },

      setDomenaHost: ({ commit }, url: string) => {
        commit('setDomenaHost', url)
      },
      setDomenaToken: ({ commit }, token: string) => {
        commit('setDomenaToken', token)
      },

      setSensiHost: ({ commit }, url: string) => {
        commit('setSensiHost', url)
      },
      setSensiToken: ({ commit }, token: string) => {
        commit('setSensiToken', token)
      },
    },
  })
}

export default store