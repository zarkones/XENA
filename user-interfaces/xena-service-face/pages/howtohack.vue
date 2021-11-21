<template>
  <div
    class = '
      ma-4
    '
  >
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
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

export default Vue.extend({
  data: () => ({
    topics: [
      {
        name: 'Cross-Site Request Forgery',
        source: 'https://en.wikipedia.org/wiki/Cross-site_request_forgery',
      },
      {
        name: 'Server-Side Request Forgery',
        source: 'https://en.wikipedia.org/wiki/Server-side_request_forgery',
      },
      {
        name: 'SQL Injections',
        source: 'https://en.wikipedia.org/wiki/SQL_injection',
      },
      {
        name: 'Command Injection',
        source: 'https://en.wikipedia.org/wiki/Code_injection',
      },
      {
        name: 'Cross-Site Scripting',
        source: 'https://en.wikipedia.org/wiki/Cross-site_scripting',
      },
      {
        name: 'Input Validation',
        source: 'https://en.wikipedia.org/wiki/Data_validation',
      },
      {
        name: 'Cross-Origin Resource Sharing',
        source: 'https://en.wikipedia.org/wiki/Cross-origin_resource_sharing',
      },
      {
        name: 'Authentication',
        source: 'https://en.wikipedia.org/wiki/Authentication',
      },
      {
        name: 'Multi-Factor Authentication',
        source: 'https://en.wikipedia.org/wiki/Multi-factor_authentication',
      },
      {
        name: 'Rate-Limiting',
        source: 'https://en.wikipedia.org/wiki/Rate_limiting',
      },
      {
        name: 'Denial of Service',
        source: 'https://en.wikipedia.org/wiki/Denial-of-service_attack',
      },
      {
        name: 'Encryption',
        source: 'https://en.wikipedia.org/wiki/Encryption',
      },
      {
        name: 'HTTP Smuggling',
        source: 'https://en.wikipedia.org/wiki/HTTP_request_smuggling',
      },
      {
        name: 'Web Cache Poisoning',
        source: 'https://en.wikipedia.org/wiki/Cache_poisoning',
      },
    ],
  }),
})
</script>

<style lang = 'css' scoped>
.topic-name {
  font-weight: bolder;
  color: #bd93f9;
  font-size: 18px;
}
</style>
