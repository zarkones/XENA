<template>
  <div>
    <v-card
      tile
      class = '
        message-card
        mb-4
      '
    >
      <v-card-text class = 'd-flex'>
        {{ message.content.shell }}

        <v-spacer></v-spacer>

        <v-btn
          x-small
          outlined
          color = 'rgba(189, 147, 249, 1)'
          @click = 'deleteMessage'
          :disabled = 'deleted'
        >
          delete
        </v-btn>
      </v-card-text>

      <div
        v-for = 'replyMessage in replies'
        :key = 'replyMessage.id'
        class = '
          pa-4
        '
      >
        {{ replyMessage }}
        <!--MessageDisplay
          :message = 'replyMessage'
        /-->
      </div>
    </v-card>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  name: 'MessageDisplay',

  components: {
  },

  data: () => ({
    draw: true,
    replies: [] as any[],
    deleted: false,
  }),

  computed: {
    ...mapGetters([
      'getPrivateKey',
      'getAtilaHost',
      'getAtilaToken',
    ])
  },

  props: {
    message: {
      required: true,
    },
  },

  methods: {
    async deleteMessage () {
      const status = await new Service.Atila(this.$axios, this.getAtilaHost, this.getAtilaToken).deleteMessage(this.message.id)
        .then(({ status }) => status)
        .catch(e => console.error(e))
      
      if (status === 204)
        this.deleted = true
    }
  },

  mounted () {
    if (this.message?.replies?.length)
      this.replies = this.message.replies
  },
})
</script>

<style lang = 'css' scoped>
.message-card {
  background-color: #44475a !important;
}
</style>
