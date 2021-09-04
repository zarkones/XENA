<template>
  <v-card>
    <v-card-content>
        <v-text-field
        dense
        v-model = 'address'
        outlined
        label = 'Address'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'submit'
      >
        Address
      </v-text-field>
    </v-card-content>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

export default Vue.extend({
  data: () => ({
    address: '',
  }),

  methods: {
    submit: () => {
      
    },
  },
})
</script>