<template>
  <v-data-table
    dense
    :headers = 'headers'
    :items = 'clients'
    item-key = 'id'
    :search = 'search'
    color = 'rgba(189, 147, 249, 1)'
    style = 'height: 100%'
    flat
    show-select
    v-model = 'selected'
    @input = 'interactionDialogUpdateClients'
  >
    <template
      v-slot:top
    >
      <v-text-field
        dense
        v-model = 'search'
        label = 'Search'
        color = 'rgba(189, 147, 249, 1)'
        class = '
          pt-4
          mx-4
        '
      ></v-text-field>
      
      <InteractionDialog />
    </template>

    <template
      v-slot:footer
    >
    </template>
  </v-data-table>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import InteractionDialog from '@/components/deshboard/interaction-dialog.vue'
import EventBus from '@/src/EventBus'

import * as Service from '@/src/services'

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

  methods: {
    interactionDialogUpdateClients () {
      EventBus.$emit('interactionDialogUpdateClients', this.selected)
    },

    async table_update (targetPlatform?: string) {
      const clients = await Service.Atila.getClients(this.$axios)
      if (clients)
        this.clients = clients
    },
  },

  mounted () {
    this.table_update()

    EventBus.$on('clients_table_update', async (targetPlatform: string) => await this.table_update(targetPlatform))
  },
})
</script>

<style lang = 'css' scoped>

</style>
