<template>
  <div>
    <iframe
      class = '
        map
      '
      frameborder = '0' 
      marginheight = '0' 
      marginwidth = '0' 
      width = '100%' 
      height = '100%' 
      scrolling = 'no'
      :src = 'mapUrl'
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
.map {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  width: 100%;
  height: 100%;
  border: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
}
</style>