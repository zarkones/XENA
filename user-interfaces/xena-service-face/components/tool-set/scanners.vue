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
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import WebFuzzer from '@/components/tools/web-fuzzer.vue'

export default Vue.extend({
  components: {
    WebFuzzer,
  },

  data: () => ({
    tools: [
      {
        name: 'Web Fuzzer',
        description: 'Web URL Fuzzer.',
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
