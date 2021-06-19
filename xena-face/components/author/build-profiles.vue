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
        Git: {{ buildProfile.gitUrl }}
        <br>
        Description: {{ buildProfile.description }}
        <br>
        Build template: {{ buildProfile.config.template }} 
        <br>
        Status: {{ buildProfile.status }}
        <br>

        <v-btn
          @click = 'downloadBuild(buildProfile.id)'
          text
          tile
          small
          color = 'rgba(189, 147, 249, 1)'
          class = '
            mt-4
          '
        >
          Download
        </v-btn>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import BuildProfiles from '@/components/author/build-profiles.vue'

import EventBus from '@/src/EventBus'

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

    downloadBuild (buildProfileId: string) {
      window.open(
        `${process.env.XENA_PYRAMID_HOST}/builds?buildProfileId=${buildProfileId}`,
        '_blank',
      )
    }
  },

  async mounted () {
    await this.getBuildProfiles()

    EventBus.$on('updateBuildProfiles', async () => await this.getBuildProfiles())
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
