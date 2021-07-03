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
      <br>
      This feature is not yet implemented. Be kind and open a Pull-Request on bitbucket.org
    </v-card-text>

    <div
      class = '
        pt-4
        mx-4
      '
    >
      <v-text-field
        v-model = 'atilaHost'
        disabled
        dense
        label = 'Atila host'
        color = 'rgba(189, 147, 249, 1)'
      ></v-text-field>

      <v-text-field
        v-model = 'pyramidHost'
        disabled
        dense
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
    atilaHost: process.env.XENA_ATILA_HOST as string,
    pyramidHost: process.env.XENA_PYRAMID_HOST as string,
  }),

  methods: {
  },

  mounted () {
  },
})
</script>

<style lang = 'css' scoped>
</style>
