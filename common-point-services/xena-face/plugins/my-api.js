

import ApiService from '~/services/Api'
/*
** Executed by ~/.nuxt/index.js with context given
** This method can be asynchronous
*/
export default ({ $axios }, inject) => {
    // Inject `api` key
    // -> app.$api
    // -> this.$api in vue components
    // -> this.$api in store actions/mutations
    inject('api', new ApiService($axios))
}