<template>
  <div>
    <v-text-field
      :disabled = '!getPrivateKey'
      dense
      v-model = 'shellCode'
      outlined
      label = 'Enter a command...'
      color = 'rgba(189, 147, 249, 1)'
      @change = 'issueMessages'
    ></v-text-field>

    <h3
      v-if = '!getPrivateKey'
    >
      Please, go to Settings -> Identity and enter your private key in order to sign the messages.
    </h3>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import jwt from 'jsonwebtoken'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    shellCode: '',
  }),

  props: {
    clients: {
      required: true,
    }
  },

  computed: {
    ...mapGetters([
      'getPrivateKey',
    ])
  },

  methods: {
    async issueMessages () {
      const message = jwt.sign(
        Buffer.from(this.shellCode).toString('base64'),
        this.getPrivateKey,
        { algorithm: 'RS512' },
      )

      const createdMessages = await Promise.all(this.clients.map(client => {
        return Service.Atila.publishMessage(this.$axios, client.id, 'shell', message)
      }))

      this.shellCode = ''
    }
  },
})
</script>

<style scoped>
</style>