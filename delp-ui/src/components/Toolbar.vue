  <template>
  <v-container class="pa-0" fluid>
    <v-toolbar elevation="1" prominent>

      <v-col cols="2" class="pa-7">
        <v-toolbar-title> <router-link style="text-decoration: none; color: inherit;" to='/'> Delp </router-link></v-toolbar-title>
      </v-col>

      <v-col cols="8">
        <SearchBar />
      </v-col>
    </v-toolbar>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import SearchBar from "../components/Search.vue";

@Component({
  components: { SearchBar },
})
export default class Toolbar extends Vue {}
</script>