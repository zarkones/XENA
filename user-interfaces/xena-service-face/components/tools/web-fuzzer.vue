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
        label = 'Request method...'
        outlined
        dense
        :loading = 'loading'
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
        :loading = 'loading'
      >
        Start
      </v-btn>

      <div
        v-for = '(response, index) in result'
        :key = 'index'
        class = '
          mt-4
        '
      >
        Url: {{ response.url }} <br>
        Status: {{ response.status }} <br>
      </div>
    </v-card-text>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    url: '',
    result: [],
    loading: false,

    wordlist: [] as string[],
    currentWord: '',
    method: '',
  }),

  computed: {
    ...mapGetters([
      'getRaHost',
      'getRaToken',
    ]),
  },

  methods: {
    async submit () {
      this.loading = true
      this.result = await new Service.Ra(this.$axios, this.getRaHost, this.getRaToken).webFuzzer(this.url, this.method, this.wordlist)
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