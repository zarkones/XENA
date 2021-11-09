<template>
  <div>
    <v-card
      tile
      class = '
        message-card
        mb-4
      '
    >
      <v-card-text>
        {{ message.content }}
      </v-card-text>

      <div
        v-for = 'replyMessage in replies'
        :key = 'replyMessage.id'
      >
        <MessageDisplay
          :message = 'replyMessage'
        />
      </div>
    </v-card>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import NodeRSA from 'node-rsa'

import { mapGetters } from 'vuex'

export default Vue.extend({
  name: 'MessageDisplay',

  components: {
  },

  data: () => ({
    draw: true,
    replies: [] as any[],
  }),

  computed: {
    ...mapGetters([
      'getPrivateKey',
    ])
  },

  props: {
    message: {
      required: true,
    }
  },

  methods: {
  },

  mounted () {
    if (this.message?.replies?.length)
      this.replies = this.message.replies.map(message => ({ ...message,
        content: new NodeRSA(this.getPrivateKey).decrypt(message.content)
      }))
  },
})
</script>

<style lang = 'css' scoped>
.message-card {
  background-color: #44475a !important;
}
</style>
