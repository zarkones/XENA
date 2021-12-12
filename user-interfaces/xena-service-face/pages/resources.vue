<template>
  <div>
    <v-data-table
    dense
    :headers = 'headers'
    :items = 'services'
    item-key = 'id'
    :search = 'search'
    color = 'rgba(189, 147, 249, 1)'
    flat
    show-select
    v-model = 'selected'
    @input = 'interactionDialogUpdateClients'
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
        class = '
          pt-4
          mx-4
        '
      ></v-text-field>
    </template>

    <!--  -->

    <template
      v-slot:footer
    >
    </template>
  </v-data-table>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { Service as ServiceEntry } from '@/src/services/Domena'
import { mapGetters } from 'vuex'

export default Vue.extend({
  components: {
  },

  async mounted () {
    await this.getServices()
  },

  computed: {
    ...mapGetters([
      'getDomenaHost',
      'getDomenaToken',
    ]),
  },

  data: () => ({
    search: '',
    services: [] as ServiceEntry[],
    headers: [
      { text: 'Address', value: 'address' },
      { text: 'Port', value: 'port' },
      { text: 'Created At', value: 'createdAt' },
    ],
  }),

  methods: {
    async getServices () {
      const services = await new Service.Domena(this.$axios, this.getDomenaHost, this.getDomenaToken).getServices()
      if (!services)
        return
      
      this.services = services
    },
  },
})
</script>