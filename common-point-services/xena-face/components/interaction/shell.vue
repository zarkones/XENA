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
      this.shellCode = ''

      console.log(createdMessages)
    }
  },
})
</script>

<style scoped>
</style>