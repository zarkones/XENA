<template>
  <v-card>
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

export default Vue.extend({
  data: () => ({
    address: '',
    result: '',
    loading: false,
  }),

  methods: {
    async submit () {
      this.loading = true
      this.result = await Service.Ra.nmap(this.$axios, this.address)
      this.loading = false
    },
  },
})
</script>