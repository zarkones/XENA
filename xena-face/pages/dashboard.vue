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
.view {
  min-width: 100% !important;
  min-height: 100% !important;
}

.zombie-view {
  min-height: 100%;
  background-color: red;
}

.tool-title {
  padding-left: 32px;
  font-size: 16px;
  font-weight: 400;
}

.panel-title {
  padding-left: 32px;
  font-size: 18px;
  font-weight: 500;
  margin-bottom: 0px !important;
}

.tool-col {
  padding-top: 0px;
}
</style>