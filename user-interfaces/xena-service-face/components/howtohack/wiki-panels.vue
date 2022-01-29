<template>
  <v-expansion-panels
    if = 'topics.length'
    accordion
    tile
  >
    <v-expansion-panel
      v-for = '(topic, topicIndex) in topics'
      :key = 'topicIndex'
    >
      <v-expansion-panel-header>
        <span
          class = '
            topic-name
          '
        >
          {{ topic.name }}
        </span>
      </v-expansion-panel-header>

      <v-expansion-panel-content>
        <iframe
          :src = 'topic.source'
          frameborder = '0'
          width = '100%'
          height = '1080'
        ></iframe>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

export default Vue.extend({
  props: {
    topics: {
      required: true,
      type: Array,
    }
  },
})
</script>

<style lang = 'css' scoped>
.topic-name {
  font-weight: bolder;
  color: #bd93f9;
  font-size: 18px;
}
</style>
