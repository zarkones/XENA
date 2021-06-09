<template>
  <v-container>
    <!--v-row
      dense    
    >
      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphLine
          title = 'All platforms'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphLine
          title = 'Windows'
          targetPlatform = 'Windows'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphLine
          title = 'Linux'
          targetPlatform = 'Linux'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphLine
          title = 'Mac'
          targetPlatform = 'MacOS'
        />
      </v-col>
    </v-row>
    
    <v-row
      dense
    >
      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphPolarArea
          title = 'All platforms'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphPolarArea
          title = 'Windows'
          targetPlatform = 'Windows'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphPolarArea
          title = 'Linux'
          targetPlatform = 'Linux'
        />
      </v-col>

      <v-col
        cols = '12'
        sm = '3'
      >
        <GraphPolarArea
          title = 'Mac'
          targetPlatform = 'MacOS'
        />
      </v-col>
    </v-row-->

    <ClientTable />
  </v-container>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import GraphClients from '@/components/analytics/graph-clients.vue'
import ClientTable from '@/components/deshboard/clients-table.vue'
import GraphPolarArea from '~/components/analytics/graph-polar.vue'
import GraphLine from '@/components/analytics/graph-line.vue'

export default Vue.extend({
  components: {
    GraphClients,
    ClientTable,
    GraphPolarArea,
    GraphLine,
  },

  mounted: () => {
  },

  data: () => ({

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