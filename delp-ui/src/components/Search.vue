<template>
  <v-container fluid>
    <v-row justify="center" align-items="start">
      <v-col cols="4">
        <v-autocomplete
          v-model="country"
          :items="countryList"
          prefix="Country"
          solo
          @change="getCities"
          autocomplete="off"
        >
          <template v-slot:no-data>
            <v-list-item>
              <v-list-item-title>
                Search for your favorite
                <strong>country</strong>
              </v-list-item-title>
            </v-list-item>
          </template></v-autocomplete
        >
      </v-col>
      <v-col cols="4">
        <v-autocomplete
          v-model="city"
          :items="cityList"
          prefix="City"
          solo
          @change="onSelectCity"
        >
          <template v-slot:no-data>
            <v-list-item>
              <v-list-item-title>
                Search for your favorite
                <strong>city</strong>
              </v-list-item-title>
            </v-list-item>
          </template></v-autocomplete
        >
      </v-col>
      <v-col cols="1">
        <v-btn dark color="primary" @click="onSelectCity" large
          ><v-icon>mdi-magnify</v-icon></v-btn
        >
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import router from "../router";

@Component
export default class Search extends Vue {
  private country = "";
  private countryList: string[] = [];

  private city = "";
  private cityList: string[] = [];

  mounted(): void {
    this.getCountries();
    if (this.country !== undefined) {
      this.getCities();
    }
  }

  private getCities() {
    console.log("fetching cities");
    fetch(`http://localhost:3000/${this.country}/cities`).then((cities) =>
      cities.json().then((data) => {
        console.log(data);
        this.cityList = data;
      })
    );
  }

  private getCountries() {
    fetch("http://localhost:3000/countries").then((countries) =>
      countries.json().then((data) => {
        console.log(data);
        this.countryList = data;
      })
    );
  }

  private onSelectCity() {
    this.$router.push({ path: `/place/${this.country}/${this.city}` });
    this.$router.go(0);
  }
}
</script>

<style>
.v-application--is-ltr .v-text-field__prefix {
  padding-right: 20px !important;
}
</style>
