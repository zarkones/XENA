<template>
  <div>
    <v-tabs
      v-model = 'selectedTab'
      style = 'color: #bd93f9 !important;caret-color: #bd93f9 !important;'
    >
      <v-tab
        v-for = 'item in tabItems'
        :key = 'item.tab'
      >
        {{ item.tab }}
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model = 'selectedTab'>
      <v-tab-item
        v-for = 'item in tabItems'
        :key = 'item.tab'
      >
        <WikiPanel
          v-if = 'item.content == "HACKING"'
          :topics = 'topics.hacking'
        />

        <WikiPanel
          v-if = 'item.content == "MALWARE"'
          :topics = 'topics.malware'
        />

        <WikiPanel
          v-if = 'item.content == "LINUX"'
          :topics = 'topics.linux'
        />

        <WikiPanel
          v-if = 'item.content == "AI"'
          :topics = 'topics.ai'
        />
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import WikiPanel from '@/components/howtohack/wiki-panels.vue'

export default Vue.extend({
  components: {
    WikiPanel,
  },
  
  data: () => ({
    selectedTab: '',
    tabItems: [
      { tab: 'Hacking', content: 'HACKING' },
      { tab: 'Malware', content: 'MALWARE' },
      { tab: 'Linux', content: 'LINUX' },
      { tab: 'Artificial Intelligence', content: 'AI' },
    ],

    topics: {
      ai: [
        {
          name: 'Artificial Intelligence',
          source: 'https://en.wikipedia.org/wiki/Artificial_intelligence',
        },
        {
          name: 'Natural Language Processing',
          source: 'https://en.wikipedia.org/wiki/Natural_language_processing',
        },
        {
          name: 'Machine Learning',
          source: 'https://en.wikipedia.org/wiki/Machine_learning',
        },
        {
          name: 'Artificial General Intelligence (AGI)',
          source: 'https://en.wikipedia.org/wiki/Artificial_general_intelligence',
        },
        {
          name: 'Existential Risk From AGI',
          source: 'https://en.wikipedia.org/wiki/Existential_risk_from_artificial_general_intelligence',
        },
        {
          name: 'Neural Networks',
          source: 'https://en.wikipedia.org/wiki/Artificial_neural_network',
        },
      ],

      linux: [
        {
          name: 'GNU/Linux',
          source: 'https://en.wikipedia.org/wiki/Linux',
        },
        {
          name: 'C Standard Library',
          source: 'https://en.wikipedia.org/wiki/C_standard_library',
        },
        {
          name: 'Linux Kernel',
          source: 'https://en.wikipedia.org/wiki/Linux_kernel',
        },
        {
          name: 'GNU Compiler Collection (GCC)',
          source: 'https://en.wikipedia.org/wiki/GNU_Compiler_Collection',
        },
        {
          name: 'Computer Architecture',
          source: 'https://en.wikipedia.org/wiki/Computer_architecture',
        },
        {
          name: 'Lodable Kernel Module',
          source: 'https://en.wikipedia.org/wiki/Loadable_kernel_module',
        },
      ],

      malware: [
        {
          name: 'Malware',
          source: 'https://en.wikipedia.org/wiki/Malware',
        },
        {
          name: 'Backdoor',
          source: 'https://en.wikipedia.org/wiki/Backdoor_(computing)',
        },
        {
          name: 'Botnet',
          source: 'https://en.wikipedia.org/wiki/Botnet',
        },
        {
          name: 'Doman Generation Algorithm (DGA)',
          source: 'https://en.wikipedia.org/wiki/Domain_generation_algorithm',
        },
        {
          name: 'Internet of Things (IoT)',
          source: 'https://en.wikipedia.org/wiki/Internet_of_things',
        },
        {
          name: 'ZeuS (Malware)',
          source: 'https://en.wikipedia.org/wiki/Zeus_(malware)',
        },
        {
          name: 'Mirai (Malware)',
          source: 'https://en.wikipedia.org/wiki/Mirai_(malware)',
        },
        {
          name: 'Stuxnet (Malware)',
          source: 'https://en.wikipedia.org/wiki/Stuxnet',
        },
      ],

      hacking: [
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
    }
  }),
})
</script>
