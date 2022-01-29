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
        @change = 'submit'
        :loading = 'loading'
      ></v-text-field>

      <p
        v-for = '(line, index) in result'
        :key = 'index'
      >
        {{ line }}
        <br>
      </p>
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
    result: [] as string[],
    loading: false,
  }),

  computed: {
    ...mapGetters([
      'getRaHost',
      'getRaToken',
    ])
  },

  methods: {
    async submit () {
      this.loading = true
      this.result = await new Service.Ra(this.$axios, this.getRaHost, this.getRaToken).sqlmap(this.url)
      this.loading = false
    },
  },
})
</script>