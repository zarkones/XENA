<template>
  <v-card
    tile
  >
    <v-card-text>
        <v-text-field
        dense
        v-model = 'url'
        outlined
        label = 'Url'
        color = 'rgba(189, 147, 249, 1)'
        :loading = 'loading'
      ></v-text-field>

      <v-select
        v-model = 'method'
        :items = 'methods'
        label = 'Request method...'
        outlined
        dense
      ></v-select>

      <v-text-field
        dense
        v-model = 'currentWord'
        outlined
        label = 'Add a word to the list...'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'addWord'
        :loading = 'loading'
      ></v-text-field>

      Wordlist: {{ wordlist.join(', ') }}

      <v-btn
        @click = 'submit'
        tile
        small
        outlined
        color = 'rgba(189, 147, 249, 1)'
        class = '
          mt-4
        '
        width = '100%'
      >
        Start
      </v-btn>

      {{ result }}
    </v-card-text>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

export default Vue.extend({
  data: () => ({
    url: '',
    result: '',
    loading: false,

    wordlist: [] as string[],
    currentWord: '',
    methods: Service.Ra.webMethods,
    method: '',
  }),

  methods: {
    async submit () {
      this.loading = true
      this.result = await Service.Ra.webFuzzer(this.$axios, this.url, this.method, this.wordlist)
      this.loading = false
    },

    addWord () {
      if (!this.currentWord)
        return
      
      this.wordlist.push(this.currentWord)
      this.currentWord = ''
    },
  },
})
</script>