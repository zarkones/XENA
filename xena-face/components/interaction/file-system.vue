<template>
  <div>
    <v-text-field
      v-if = ctx
      :label = ctx.location
      :loading = loading
      color = 'purple accent-2'
      dense
      v-model = new_path
      @change = "$emit('direct-relocation', new_path); new_path = ''"
    ></v-text-field>

    <v-simple-table
      dense
    >
      <template v-slot:default>
        <thead>
          <tr>
            <th
              class = '
                text-left
              '
            >
              Name
            </th>

            <th
              class = '
                text-left
              '
            >
              Actions
            </th>

            <!--th
              class = '
                text-left
              '
            >
              Size (MB)
            </th>

            <th
              class = '
                text-left
              '
            >
              Executable
            </th-->
          </tr>
        </thead>
        <tbody
          v-if = ctx
        >
          <tr 
            v-for = '(file, key) in ctx.files' 
            :key = key
          >
            <td>
              {{ file }}
            </td>

            <td>
              <v-btn
                x-small
                depressed
                class = '
                  purple darken-2
                '
                @click= "$emit('onClick', file)"
              >
                Enter
              </v-btn>
            </td>
            <!--td>{{ file.name }}</td>
            <td>{{ file.size }}</td>
            <td
              v-if = file.executable
              class = 'blue--text'
            >
              {{ file.executable }}
            </td>
            <td
              v-if = !file.executable
              class = 'grey--text'
            >
              {{ file.executable }}
            </td-->
          </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  components: {
  },

  props: {
    loading: Boolean,
    ctx: {
      required: true,
    }
  },

  data () {
    return {
      new_path: ''
    }
  }
}
</script>

<style lang = 'css'>
.panel-title {
  padding-left: 32px;
  font-size: 18px;
  font-weight: 500;
  margin-bottom: 0px !important;
}
</style>