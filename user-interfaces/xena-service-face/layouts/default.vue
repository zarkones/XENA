<template>
  <v-app dark>
    <v-navigation-drawer
      v-model = 'drawer'
      :mini-variant = 'miniVariant'
      :clipped = 'clipped'
      fixed
      app
      expand-on-hover
    >
      <v-list>
        <v-list-item
          v-for = '(item, i) in items'
          :key = 'i'
          :to = 'item.to'
          router
          exact
        >
          <v-list-item-action>
            <v-icon
              v-if = '!item.text'
            >
              {{ item.icon }}
            </v-icon>
            <h5
              v-if = 'item.text'
              class = '
                pa-1
              '
            >
              {{ item.text }}
            </h5>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text = 'item.title' />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-system-bar
      class = '
        system-bar
      '
    >
      <ChatWithXena />

      <v-spacer></v-spacer>

      <v-btn
        small
        text
        light
      >
        {{ getUsername ? getUsername : 'Not logged in' }}
      </v-btn>
    </v-system-bar>

    <!-- Pages -->
    <v-main>
      <nuxt />
    </v-main>

    <!--v-footer
      :absolute = '!fixed'
      app
      outlined
      fixed
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
      <v-btn
        tile
        small
        color = 'rgba(189, 147, 249, 1)'
        :plain = 'true'
        v-if = 'getUsername'
      >
        {{ getUsername }}
      </v-btn>

    </v-footer-->
  </v-app>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import { mapGetters } from 'vuex'

import ChatWithXena from '@/components/xena-chat/chat-dialog.vue'

export default Vue.extend({
  components: {
    ChatWithXena,
  },

  computed: {
    ...mapGetters([
      'getPrivateKey',
      'getUsername',
    ]),
  },

  mounted () {
    if (this.getPrivateKey && this.getUsername && this.$router.history.current.path == '/')
      this.$router.push('/dashboard')
  },

  data: () => ({
    searchTerm: '',

    clipped: false,
    drawer: true,
    miniVariant: false,
    right: true,
    rightDrawer: false,
    title: 'XENA',
    fixed: true,

    items: [
      {
        icon: 'mdi-view-dashboard',
        title: 'Dashboard',
        to: '/dashboard'
      },
      {
        icon: 'mdi-console',
        title: 'Control Terminal',
        to: '/terminal'
      },
      {
        icon: 'mdi-alphabet-piqad',
        title: 'Hacking Tools',
        to: '/tools',
      },
      {
        icon: 'mdi-book',
        title: 'Hacking Manual',
        to: '/howtohack',
      },
      {
        icon: 'mdi-cube-outline',
        title: 'Cloud Build',
        to: '/build'
      },
      // This feature isn't ready yet.
      // But it's worth exploring.
      // {
      //   icon: 'mdi-forum',
      //   title: 'Forum',
      //   to: '/forum',
      // },
      {
        icon: 'mdi-web',
        title: 'Internet Registry',
        to: '/resources',
      },
      // {
      //   icon: 'mdi-car',
      //   title: 'Vehicle Registry',
      //   to: '/registry',
      // },
      {
        icon: 'mdi-earth',
        title: 'Map',
        to: '/map',
      },
      {
        icon: 'mdi-wrench',
        title: 'Settings',
        to: '/settings'
      },
    ],
  }),
})

</script>

<style scoped>
.system-bar {
  background-color: #bd93f9;
  bottom: 0px !important;
  position: absolute;
  min-width: 100% !important;
  z-index: 100;
}
</style>

<style>
.logo {
  color: white;
}

.graph-number {
  padding-left: 12px;
  color: #6272a4;
  font-weight: 500;
}

.purple-color {
  color: #bd93f9 !important;
}
.green-color {
  color: #50fa7b !important;
}
.red-color {
  color: #ff5555 !important;
}
.pink-color {
  color: #ff79c6 !important;
}
.blue-color {
  color: #8be9fd !important;
}
.white-color {
  color: #f8f8f2 !important;
}
.dark-color {
  color: #44475a !important;
}
.darker-color {
  color: #282a36 !important;
}

.bg-dark-color {
  background-color: #44475a !important;
}

.bg-darker-color {
  background-color: #282a36 !important;
}


/* 
  Personalization of vuetify.
 */
.v-sheet.v-card:not(.v-sheet--outlined) {
  box-shadow: none !important;
}
.v-expansion-panel::before {
  box-shadow: none !important;
}
.v-data-table {
  background-color: #282a36 !important;
}
.v-expansion-panel-header > *:not(.v-expansion-panel-header__icon) {
  flex: none !important;
}
.v-expansion-panel {
  background-color: #44475a !important;
}
.v-sheet {
  background-color: #282a36 !important;
}
.theme--dark.v-application {
  background-color: #282a36 !important;
}
.theme--dark.v-card {
  background-color: #282a36 !important;
}
.theme--dark.v-window {
  background-color: #282a36 !important;
}
.theme--dark.v-slide-group {
  background-color: #282a36 !important;
}
.theme--dark.v-navigation-drawer {
  background-color: #282a36 !important;
}
.v-list .v-list-item--active {
  color: #bd93f9 !important;
}
.theme--dark.v-list-item:not(.v-list-item--active):not(.v-list-item--disabled) {
  color: #bd93f9 !important;
}
.v-icon {
  color: #bd93f9 !important;
}
.theme--light.v-btn {
  color: #282a36 !important;
}
.theme--dark.v-sheet.v-list {
  padding-top: 0px;
}
</style>
