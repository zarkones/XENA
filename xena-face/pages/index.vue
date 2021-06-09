<template>
  <v-container
    fill-height
    fluid
  >
    <v-row
      align = 'center'
      justify = 'center'
    >
      <v-col>
        <v-card
          max-width = '800px'
          class = '
            mx-auto
          '
          justify = 'center'
          outlined
          tile
        >
          <v-card-title>
            Please, login.
          </v-card-title>

          <v-divider></v-divider>

          <v-text-field
            dense
            v-model = 'username'
            label = 'Username'
            color = 'rgba(189, 147, 249, 1)'
            class = '
              pt-4
              mx-4
            '
          ></v-text-field>

          <v-text-field
            dense
            type = 'password'
            v-model = 'password'
            label = 'Password'
            color = 'rgba(189, 147, 249, 1)'
            class = '
              pt-4
              mx-4
            '
          ></v-text-field>

          <v-btn
            block
            text
            tile
            small
            color = 'rgba(189, 147, 249, 1)'
            @click = 'login'
          >
            Login
          </v-btn>
        </v-card>
      </v-col>
    </v-row>

    <v-snackbar
      v-model = 'auth_msg'
      :timeout = '2000'
    >
      {{ auth_msg_payload }}

      <template v-slot:action = '{ attrs }'>
        <v-btn
          color='rgba(189, 147, 249, 1)'
          text
          v-bind = 'attrs'
          @click = 'auth_msg = false'
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script lang = 'ts'>
import Vue from 'vue'
import AuthRepo from '@/src/api/auth'

export default Vue.extend({
  data: () => ({
    username: '',
    password: '',

    password_start: 0,

    auth_msg: false,
    auth_msg_payload: 'Wrong username or password.',
  }),

  methods: {
    async login () {
      // const response = await AuthRepo.login(this.$axios, this.username, this.password)
      // if (!response) {
      //   this.auth_msg = true
      //   return
      // }

      this.$auth.loginWith('local', {
        username: this.username,
        password: this.password,
      })

      this.$router.push('/analytics')
    },
  },
})
</script>
