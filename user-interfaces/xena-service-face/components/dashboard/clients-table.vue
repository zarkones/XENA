<template>
  <v-data-table
    dense
    :headers = 'headers'
    :items = 'clients'
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
</template>

<script lang = 'ts'>
import Vue from 'vue'

import InteractionDialog from '@/components/dashboard/interaction-dialog.vue'

import EventBus from '@/src/EventBus'
import * as Service from '@/src/services'

import { mapGetters } from 'vuex'

export default Vue.extend({
  components: {
    InteractionDialog,
  },

  data: () => ({
    search: '',
    selected: [],
    clients: [] as any[],
    dialog: false,
    headers: [
      { text: 'ID', value: 'id' },
    ]
  }),

  computed: {
    ...mapGetters([
      'getAtilaHost',
    ])
  },

  methods: {
    interactionDialogUpdateClients () {
      EventBus.$emit('interactionDialogUpdateClients', this.selected)
      if (this.selected.length)
        EventBus.$emit(`interactionDialogUpdateSelectedClient`, this.selected[0].id)
    },

    async tableUpdate (targetPlatform?: string) {
      const clients = await new Service.Atila(this.$axios, this.getAtilaHost).getClients()
      if (clients)
        this.clients = clients
    },
  },

  mounted () {
    this.tableUpdate()

    EventBus.$on('clientsTableUpdate', async (targetPlatform: string) => await this.tableUpdate(targetPlatform))
  },
})
</script>

<style lang = 'css' scoped>
</style>
