<template>
  <div
    class = '
      ma-4
    '
  >
    <v-row
      dense    
    >
      <v-col
        cols = '12'
        sm = '6'
      >
        <GraphLine
          title = 'All Platforms'
        />
      </v-col>
    
      <v-col
        cols = '12'
        sm = '6'
      >
        <GraphPolarArea
          title = 'All platforms'
        />
      </v-col>
    </v-row>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import GraphPolarArea from '~/components/analytics/graph-polar.vue'
import GraphLine from '@/components/analytics/graph-line.vue'

export default Vue.extend({
  components: {
    GraphPolarArea,
    GraphLine,
  },

  mounted () {
  },

  methods: {
  },

  data: () => ({
  }),
})
</script>

<style lang = 'css' scoped>
</style>