<template>
  <v-app dark>
    <v-navigation-drawer
      v-model = 'drawer'
      :mini-variant = 'miniVariant'
      :clipped = 'clipped'
      fixed
      app
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

    <!-- Pages -->
    <v-main>
      <nuxt />
    </v-main>

    <!--v-footer
      :absolute="!fixed"
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer-->
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      clipped: false,
      drawer: true,
      fixed: false,
      items: [
        {
          icon: 'mdi-console',
          title: 'Dashboard',
          to: '/dashboard'
        },
        {
          icon: 'mdi-cube-outline',
          title: 'Author Studio',
          to: '/author'
        },
        {
          icon: 'mdi-wrench',
          title: 'Settings',
          to: '/settings'
        },
      ],
      miniVariant: true,
      right: true,
      rightDrawer: false,
      title: 'XENA'
    }
  }
}

</script>

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
  color: #6272a4 !important;
}
.v-icon {
  color: #bd93f9 !important;
}
.theme--light.v-btn {
  color: #282a36 !important;
}
</style>
