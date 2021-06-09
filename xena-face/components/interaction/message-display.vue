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

export default Vue.extend({
  name: 'MessageDisplay',

  components: {
  },

  data: () => ({
    draw: true,
    replies: [] as any[],
  }),

  props: {
    message: {
      required: true,
    }
  },

  methods: {
  },

  mounted () {
    if (this.message?.replies?.length)
      this.replies = this.message.replies.map(message => ({ ...message, content: Buffer.from(message.content, 'base64') }))
  },
})
</script>

<style lang = 'css' scoped>
.message-card {
  background-color: #44475a !important;
}
</style>
