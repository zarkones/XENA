<template>
  <v-card
    class = '
      mx-auto
      text-center
    '
    outlined
    tile
  >
    <v-btn
      block
      text
      disabled
      tile
      small
      color = 'rgba(189, 147, 249, 1)'
    >
      {{ title }}
    </v-btn>

    <v-divider></v-divider>

    <canvas
      :id = '`line-chart-${targetPlatform}`'
    ></canvas>
  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import { mapGetters } from 'vuex'
import { Chart } from 'chart.js'

import * as Service from '@/src/services'

export default Vue.extend({
  data: () => ({
    botClients: [] as number[],
    chart: null,
    start: '',
    end: '',
  }),

  props: {
    targetPlatform: {
      required: false,
      type: String,
    },

    title: {
      required: true,
      type: String,
    }
  },

  methods: {
    chartInit (update?: boolean) {
      if (this.chart && !update)
        return

      this.chart = new Chart(document.getElementById(`line-chart-${this.targetPlatform}`), {
        type: 'line',
        data: {
          labels: this.botClients,
          datasets: [
            { 
              data: this.botClients,
              label: 'Clients',
              pointBackgroundColor: '#bd93f9',
              pointBorderColor: '#bd93f9',
              borderColor: '#6272a4',
              borderWidth: 2,
              radius: 1.5,
              fill: true,
            }
          ]
        },
        options: {
          title: {
            display: false,
            text: this.title
          },
          scales: {
            xAxes: [{
              ticks: {
                display: false
              },
              gridLines: {
                display: false
              }
            }],
            yAxes: [{
              ticks: {
                display: false
              },
              gridLines: {
                display: true
              }
            }]
          },
          legend: {
            display: false
          },
          tooltips: {
            enabled: true
          },
        }
      })
    },

    async getStatsCount () {
      const count = await new Service.Atila(this.$axios, this.getAtilaHost, this.getAtilaToken).getCount()
      if (count?.length)
        this.botClients = count
    },
  },

  computed: {
    ...mapGetters([
      'getAtilaHost',
      'getAtilaToken',
    ])
  },
  
  async mounted () {
    await this.getStatsCount()

    this.chartInit()
  }
})
</script>

<style lang = 'css' scoped>

</style>