<template>
  <v-expansion-panels
    if = 'buildProfiles.length'
    accordion
  >
    <v-expansion-panel
      v-for = '(buildProfile, buildProfileIndex) in builldProfiles'
      :key = 'buildProfileIndex'
    >
      <v-expansion-panel-header>
        <span
          class = '
            build-profile-name
          '
        >
          {{ buildProfile.name }}
        </span>
      </v-expansion-panel-header>

      <v-expansion-panel-content>
        Gir: {{ buildProfile.gitUrl }}
        <br>
        Status: {{ buildProfile.status }}
        <br>
        Description: {{ buildProfile.description }}
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import BuildProfiles from '@/components/author/build-profiles.vue'

import * as Service from '@/src/services'

export default Vue.extend({
  components: {
    BuildProfiles,
  },

  data: () => ({
    builldProfiles: [] as any[]
  }),

  methods: {
    async getBuildProfiles () {
      const builldProfiles = await Service.Pyramid.getBuilldProfiles(this.$axios)
      if (builldProfiles)
        this.builldProfiles = builldProfiles
    },
  },

  async mounted () {
    await this.getBuildProfiles()
  },
})
</script>

<style lang = 'css' scoped>
.build-profile-name {
  font-weight: bolder;
  color: #bd93f9;
  font-size: 18px;
}
</style>
