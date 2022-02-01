<template>
  <v-dialog
    v-model = 'dialog'
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
          @click = 'dialog = false'
        >
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>

      <div
        v-for = 'message in messages'
        :key = 'message.id'
        class = '
          ma-4
        '
      >
        <span
          class = '
            input-display
          '
        >
          {{ message.input }}
        </span>
        <br>
        {{ message.output }}
        <br>
      </div>

      <v-card-actions>
        <v-text-field
          dense
          v-model = 'input'
          outlined
          label = 'Send a message...'
          color = 'rgba(189, 147, 249, 1)'
          @change = 'issueMessage'
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

export default Vue.extend({
  data: () => ({
    dialog: false,
    input: '',
    messages: [
      {
        input: 'What is malware?',
        output: 'Malware is a malicious software designed to perform an action or a set of actions which are undesired by the affected user.',
      },
      {
        input: 'Is C++ a language?',
        output: 'Correct. C++ is a low level programming language.'
      }
    ],
  }),

  methods: {
    async issueMessage () {

    },

    async fetchMessages () {

    }
  },

  mounted () {
    this.fetchMessages()
  },
})

</script>

<style scoped>
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