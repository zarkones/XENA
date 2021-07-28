<template>
  <v-text-field
      dense
      v-model = 'shellCode'
      outlined
      label = 'Enter a command...'
      color = 'rgba(189, 147, 249, 1)'
      @change = 'issueMessages'
    ></v-text-field>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import jwt from 'jsonwebtoken'

export default Vue.extend({
  data: () => ({
    shellCode: '',
  }),

  props: {
    clients: {
      required: true,
    }
  },

  methods: {
    async issueMessages () {
      const createdMessages = await Promise.all(this.clients.map(client => {
        return Service.Atila.publishMessage(this.$axios, client.id, 'shell', Buffer.from(this.shellCode).toString('base64'))
      }))

      console.log(jwt.sign({ foo: 'bar' }, `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`, { algorithm: 'RS256'}))

      this.shellCode = ''

      console.log(createdMessages)
    }
  },
})
</script>

<style scoped>
</style>