<template>
  <v-row>
    <v-col>
      <ChartMain />
    </v-col>
    <v-col>
      <ClientTable />
    </v-col>
  </v-row>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import ChartMain from '@/components/analytics/graph-main.vue'
import ClientTable from '@/components/deshboard/clients-table.vue'

export default Vue.extend({
  components: {
    ClientTable,
    ChartMain,
  },
})
</script>