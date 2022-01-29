<template>
  <v-card
    tile
  >
    <v-card-text>
      <v-text-field
        dense
        v-model = 'domain'
        outlined
        label = 'Domain'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'submit'
        :loading = 'loading'
      ></v-text-field>

      <v-text-field
        dense
        v-model = 'rawDict'
        outlined
        label = 'Comma separated list of subdomain names...'
        color = 'rgba(189, 147, 249, 1)'
        @change = 'submit'
        :loading = 'loading'
      ></v-text-field>
        
      <div
        v-if = 'result && result.alive && result.dead'
      >
        Alive:
        <p
          v-for = '(name, index) in result.alive'
          :key = 'index'
        >
          {{ name }}
        </p>

        <br>

        Dead:
        <p
          v-for = '(name, index) in result.dead'
          :key = 'index'
        >
          {{ name }}
        </p>
      </div>

      

    </v-card-text>
    

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { BruteForcedSubdomains } from '@/src/services/Ra' 
import { mapGetters } from 'vuex'

export default Vue.extend({
  data: () => ({
    result: {} as BruteForcedSubdomains,
    domain: '',
    rawDict: '',
    loading: false
  }),
  
  computed: {
    ...mapGetters([
      'getRaHost',
      'getRaToken',
    ]),
  },

  methods: {
    async submit () { 
      if (this.domain && (this.rawDict && this.rawDirct?.length)) {
        this.loading = true
        const result = await new Service.Ra(this.$axios, this.getRaHost, this.getRaToken).subdomainBruteforce(this.domain, this.rawDict.split(','))
        if (!result)
          return
        
        this.result = result
        this.loading = false
      }
    },
  },
})
</script>