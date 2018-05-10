<template>
  <v-container fluid>
      <v-card >
    <template v-for="spider in spiders" >
        <v-list three-line subheader :key="spider.Name">
          <v-subheader>{{spider.Name}}</v-subheader>
          <template v-for="(item, index) in spider.Out">
            <v-list-tile
              @click="toggle(item)"
              :key="item.Title"
            >
              <v-list-tile-content>
                <v-list-tile-title>{{ item.Title }}</v-list-tile-title>
                <v-list-tile-sub-title>{{ item.Subtitle }}</v-list-tile-sub-title>
              </v-list-tile-content>
              <v-list-tile-action>
                <v-list-tile-action-text>{{ item.Stars }}</v-list-tile-action-text>
              </v-list-tile-action>
            </v-list-tile>
            <v-divider v-if="index + 1 < spider.Out.length" :key="index"></v-divider>
          </template>
        </v-list>
    </template>
      </v-card>
  </v-container>
</template>

<script>
import { getData } from '@/api/table'
// import {axios} from 'axios'

export default {
  data () {
    return {
      spiders: []
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      getData('/api/hotspots').then(response => {
        // console.log(response)
        this.spiders = response.data.data.Data
        console.log(this.spiders)
      })
    },
    toggle (item) {
      window.open(item.URL)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
  padding: 0px
}
.list__tile {
    height: auto;
}
.list__tile__sub-title {
  white-space: normal
}
/* .list__tile__title {
  white-space: normal
} */
</style>
