<template>
  <v-data-table
    dense
    :headers = 'headers'
    :items = 'vehicles'
    item-key = 'id'
    :search = 'search'
    color = 'rgba(189, 147, 249, 1)'
    flat
    show-select
    v-model = 'selected'
    :single-select = 'false'
  >
    <template
      v-slot:top
    >
      <v-text-field
        outlined
        dense
        v-model = 'search'
        label = 'Search'
        color = 'rgba(189, 147, 249, 1)'
        class = '
          pt-4
          mx-4
        '
      ></v-text-field>
    </template>

    <!--  -->

    <template
      v-slot:footer
    >
    </template>
  </v-data-table>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import InteractionDialog from '@/components/dashboard/interaction-dialog.vue'

type VehicleType =
  | 'MICRO'
  | 'SEDAN'
  | 'CUV'
  | 'SUV'
  | 'HATCHBACK'
  | 'ROEADSTER'
  | 'PICKUP'
  | 'VAN'
  | 'COUPE'
  | 'SUPERCAR'
  | 'CAMPERVAN'
  | 'MINI_TRUCK'
  | 'CABRIOLET' 
  | 'MINIVAN'
  | 'TRUCK'
  | 'BIG_TRUCK'

type Vehicle = {
  id: string
  plate: string
  type: VehicleType
}

export default Vue.extend({
  components: {
    InteractionDialog,
  },

  data: () => ({
    search: '',
    selected: [],
    vehicles: [
      {
        id: 'asdasd-asdasdasd-asdasd-asdasdas',
        plate: 'GH-546-747',
        type: 'SEDAN',
      },
      {
        id: 'werwetw-ghjghj-xcvwetsfdt-dtughikh',
        plate: 'AF-365-124',
        type: 'BIG_TRUCK',
      },
    ] as Vehicle[],
    dialog: false,
    headers: [
      { text: 'Plate', value: 'plate' },
      { text: 'Type', value: 'type' },
    ]
  }),

  methods: {
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
</style>
