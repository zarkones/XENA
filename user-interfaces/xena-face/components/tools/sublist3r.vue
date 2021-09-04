<template>
  <v-card>
    <v-card-text>
        <v-text-field
        dense
        v-model = 'domain'
        outlined
        label = 'Domain'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'submit'
        :loading = 'loading'
      >
        Domain
      </v-text-field>

      <div
        v-if = 'subdomains && subdomains.length'
      >
        <p
          v-for = '(name, index) in subdomains'
          :key = 'index'
        >
          {{ name }}
        </p>
      </div>

    </v-card-text>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

export default Vue.extend({
  data: () => ({
    domain: '',
    subdomains: [] as string[],
    loading: false,
  }),

  methods: {
    async submit () {
      this.loading = true
      this.subdomains = await Service.Ra.sublist3r(this.$axios, this.domain)
      this.loading = false
    },
  },
})
</script>