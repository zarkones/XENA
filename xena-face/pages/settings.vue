<template>
  <v-card
    tile
    flat
  >
    <v-card-title>
      Settings
    </v-card-title>

    <v-card-text>
      Here you'll be able to configure your client's settings.
    </v-card-text>

    <div
      class = '
        pt-4
        mx-4
      '
    >
      <v-text-field
        disabled
        dense
        v-model = 'atilaHost'
        label = 'Atila host'
        color = 'rgba(189, 147, 249, 1)'
      ></v-text-field>

      <v-text-field
        disabled
        dense
        v-model = 'pyramidHost'
        label = 'Pyramid host'
        color = 'rgba(189, 147, 249, 1)'
      ></v-text-field>
    </div>

  </v-card>
</template>

<script lang = 'ts'>
import Vue from 'vue'

export default Vue.extend({
  components: {
  },

  data: () => ({
    atilaHost: '',
    pyramidHost: '',
  }),

  methods: {
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
</style>
