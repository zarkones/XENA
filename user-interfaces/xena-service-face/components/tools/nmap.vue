<template>
  <v-card
    tile
  >
    <v-card-text>
        <v-text-field
        dense
        v-model = 'address'
        outlined
        label = 'Address'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'submit'
        :loading = 'loading'
      >
        Address
      </v-text-field>

      {{ result }}
    </v-card-text>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    address: '',
    result: '',
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
      this.result = await new Service.Ra(this.$axios, this.getRaHost, this.getRaToken).nmap(this.address)
      this.loading = false
    },
  },
})
</script>