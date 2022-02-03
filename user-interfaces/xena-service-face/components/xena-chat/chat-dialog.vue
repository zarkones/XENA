<template>
  <v-dialog
    v-model = 'popup'
    fullscreen
    hide-overlay
    transition = 'dialog-bottom-transition'
  >
    <template v-slot:activator = '{ on, attrs }'>
      <v-btn
        v-bind = 'attrs'
        v-on = 'on'
        small
        text
        light
      >
        Chat With XENA
      </v-btn>
    </template>

    <v-card>
      <v-toolbar
        dense
        tile
        flat
      >
        <v-toolbar-title></v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn
          icon
          dark
          @click = 'popup = false'
        >
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>

      <div
        class = '
          sentences
        '
      >
        <div
          v-for = 'sentence in dialog'
          :key = 'sentence.id'
          class = '
            ma-4
          '
        >
          <span
            class = '
              input-display
            '
          >
            {{ sentence.input }}
          </span>
          <br>
          {{ sentence.output }}
          <br>
        </div>
      </div>

      <v-card-actions>
        <v-text-field
          dense
          v-model = 'input'
          outlined
          label = 'Send a message...'
          color = 'rgba(189, 147, 249, 1)'
          @change = 'issueMessage'
          :loading = 'loading'
          class = '
            input-field
            pl-4
            pr-4
          '
        ></v-text-field>
      </v-card-actions>
    </v-card>

  </v-dialog>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'
import { Dialog } from '@/src/services/Sensi'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    popup: false,
    input: '',
    loading: false,
    dialog: [] as Dialog,
  }),

  computed: {
    ...mapGetters([
      'getSensiHost',
      'getSensiToken',
    ])
  },

  methods: {
    async issueMessage () {
      if (!this.input)
        return
      
      const input = this.input
      this.input = ''

      this.loading = true
      await new Service.Sensi(this.$axios, this.getSensiHost, this.getSensiToken).insert(input)
      this.loading = false

      await this.fetchMessages()
    },

    async fetchMessages () {
      const dialog = await new Service.Sensi(this.$axios, this.getSensiHost, this.getSensiToken).getDialog()
      if (!dialog)
        return
      
      this.dialog = dialog
    }
  },

  mounted () {
    this.fetchMessages()
  },
})

</script>

<style scoped>
.sentences {
  padding-bottom: 42px;
}

.input-display {
  color: rgba(189, 147, 249, 1);
}

.input-field {
  position: fixed;
  bottom: 0px !important;
  min-width: 100%;
  right: 0px;
}
</style>