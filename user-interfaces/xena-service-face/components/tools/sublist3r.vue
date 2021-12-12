<template>
  <v-card
    tile
  >
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

      <!--div
        v-if = 'subdomains && subdomains.length'
      >
        <p
          v-for = '(name, index) in subdomains'
          :key = 'index'
        >
          {{ name }}
        </p>
      </div-->

    <v-data-table
      dense
      :headers = 'headers'
      :items = 'subdomains'
      item-key = 'name'
      :search = 'search'
      color = 'rgba(189, 147, 249, 1)'
      flat
      show-select
      v-model = 'selected'
      :single-select = 'true'
    >
      <template
        v-slot:top
      >
        <v-text-field
          outlined
          dense
          v-model = 'search'
          label = 'Search'
          color = 'rgba(189, 147, 249, 1)'
        ></v-text-field>
      </template>

      <!--  -->

      <template
        v-slot:footer
      >
      </template>
    </v-data-table>

    </v-card-text>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'
import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    domain: '',
    subdomains: [] as string[],
    loading: false,
    selected: [] as string[],

    search: '',

    headers: [
      { text: 'name', value: 'name' },
    ]
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
      const subdomains = await new Service.Ra(this.$axios, this.getRaHost, this.getRaToken).sublist3r(this.domain)
      this.loading = false

      if (subdomains)
        this.subdomains = subdomains.map(name => ({ name }))
    },
  },
})
</script>