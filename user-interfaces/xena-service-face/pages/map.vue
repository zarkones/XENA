<template>
  <div>
    <iframe
      width = '1920'
      height = '1080'
      frameborder = '0'
      scrolling = 'no'
      marginheight = '0'
      marginwidth = '0'
      :src = 'mapUrl'
      style = 'border: 1px solid black'
    ></iframe>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

export default Vue.extend({
  components: {
  },

  mounted: () => {
  },

  data: () => ({
    mapUrl: 'https://www.openstreetmap.org/export/embed.html?layer=opnvkarte',
  }),
})
</script>

<style lang = 'css'>
</style>