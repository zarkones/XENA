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
        {{ message.content.shell }}
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
    },
  },

  methods: {
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
