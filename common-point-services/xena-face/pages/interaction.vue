<template>
  <div>
    <InteractionDialog 
      style = '
        margin-right: 408px;
      '
    />

    <v-navigation-drawer
      v-model = 'clientsTable'
      :right = 'true'
      fixed
      :clipped = 'true'
      :mini-variant = 'false'
      width = '408'
    >
      <ClientsTable />
    </v-navigation-drawer>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import InteractionDialog from '@/components/dashboard/interaction-dialog.vue'
import ClientsTable from '@/components/dashboard/clients-table.vue'

export default Vue.extend({
  components: {
    ClientsTable,
    InteractionDialog,
  },

  mounted: () => {
  },

  data: () => ({
    clientsTable: true,
  }),
})
</script>

<style lang = 'css'>
</style>