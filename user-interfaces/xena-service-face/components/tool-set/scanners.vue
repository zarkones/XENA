<template>
  <v-expansion-panels
    if = 'tools.length'
    accordion
    tile
  >
    <v-expansion-panel
      v-for = '(tool, toolIndex) in tools'
      :key = 'toolIndex'
    >
      <v-expansion-panel-header>
        {{ tool.name }}
      </v-expansion-panel-header>

      <v-expansion-panel-content>
        <WebFuzzer
          v-if = 'tool.name == "Web Fuzzer"'
        />
        <SqlMap
          v-if = 'tool.name == "SqlMap"'
        />
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import WebFuzzer from '@/components/tools/web-fuzzer.vue'
import SqlMap from '@/components/tools/sqlmap.vue'

export default Vue.extend({
  components: {
    WebFuzzer,
    SqlMap,
  },

  data: () => ({
    tools: [
      {
        name: 'Web Fuzzer',
        description: 'Web URL Fuzzer.',
      },
      {
        name: 'SqlMap',
        description: 'Database enumeration tool.',
      },
    ] as const
  }),

  methods: {
  },

  mounted () {
  },
})
</script>

<style scoped>
.v-application p {
  margin-bottom: 0px !important;
}

.v-text-field.v-text-field--enclosed .v-text-field__details {
  margin-bottom: 0px !important;
}
</style>
