<template>
  <v-card
    flat
    tile
    min-height = '190px'
    @click = 'setBuildTemplate()'
    :disabled = 'details.disabled'
  >
    <v-card-title>
      {{ details.name }}

      <v-spacer></v-spacer>

      <span
        v-if = 'details.disabled'
      >
        [cooming soon]
      </span>
    </v-card-title>
  
    <v-divider></v-divider>

    <v-row
      no-gutters
    >
      <v-col
        cols = '12'
        sm = '2'
      >
        <v-img
          :src = 'details.logo'
          max-height = '126px'
          max-width = '126px'
          contain
        ></v-img>
      </v-col>

      <v-col>
        <v-card-text>
          {{ details.details }}
        </v-card-text>
      </v-col>
    </v-row>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import EventBus from '@/src/EventBus'

export default Vue.extend({
  data: () => ({
    isActive: false,
  }),

  props: {
    details: {
      type: Object,
      required: true,
    }
  },

  methods: {
    setBuildTemplate () {
      EventBus.$emit('updateBuildTemplate', this.details.id)
    }
  },
})
</script>
