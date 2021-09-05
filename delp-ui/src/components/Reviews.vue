<template>
  <v-container style="height: 100vh;">

    <h1 class="text-center"> <span class="text-h3">What people say about </span>  <span class="text-h2">Kuala Lumpur</span> </h1> 
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vitae dictum nulla. Quisque pulvinar dolor eget sem interdum vulputate. Donec tristique molestie mattis. Mauris non commodo ligula, vitae dignissim lectus. Nulla ullamcorper, massa ac porttitor tincidunt, nisl arcu auctor orci, at sodales lorem tortor non elit. Proin consectetur porta elit vel lacinia. Integer varius lacus eget euismod convallis. Sed quis quam a libero congue placerat ut sed nibh. Donec a volutpat quam.

Nulla in felis ut ligula lobortis congue. Cras sollicitudin enim posuere ipsum gravida, non maximus turpis auctor. Maecenas aliquam vitae nisi quis suscipit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut orci eros, pellentesque et pretium ut, tincidunt quis velit. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Suspendisse placerat dapibus vulputate. Mauris id massa quis quam malesuada auctor. Proin dictum felis nisi, eget rutrum elit condimentum et. Morbi ut sem in orci porttitor gravida at quis nisi. Phasellus scelerisque eros sit amet nulla egestas, nec maximus leo imperdiet. Vivamus vehicula eget lectus in sodales.

Aliquam sit amet lobortis dui, ac aliquet augue. Etiam sit amet elit cursus est vehicula commodo. Aliquam erat volutpat. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras sit amet velit faucibus dui tempor interdum ac maximus mauris. Fusce sagittis vel quam vitae eleifend. Nulla facilisis nibh ut nunc commodo venenatis. In hendrerit nibh sit amet felis efficitur, eu pulvinar tellus cursus. Nulla malesuada felis ut massa tristique, sit amet convallis nisi scelerisque. Fusce eu justo eu dui aliquam molestie ac quis nisi. Nulla vel turpis et lacus varius dictum at eget orci. Integer ac risus nisi. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Nullam mollis non felis quis lobortis.

Ut a nunc id risus fermentum sagittis. Vestibulum non sem diam. Nam ac pulvinar turpis. Vestibulum nisi purus, porta et purus sed, malesuada consectetur diam. Suspendisse sit amet mollis lacus, et volutpat turpis. Vivamus non lacinia leo, eu vehicula libero. Morbi lacinia risus a purus pharetra pulvinar. Vivamus ultricies mi vulputate massa tempor cursus.

Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec id lobortis ipsum. In laoreet urna sapien, ut viverra orci commodo et. In eget neque lectus. Curabitur eu tellus ut leo convallis maximus. Pellentesque vulputate interdum blandit. Duis fermentum volutpat gravida. Mauris neque nulla, ultricies et magna euismod, ultrices laoreet mi. Nulla facilisi. Quisque quis placerat ante. Praesent non congue erat. Donec pellentesque molestie pretium. Aenean pharetra molestie nibh, nec aliquam lectus tempor ut. In consequat magna nulla, non eleifend libero hendrerit non.
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component({ components: { } })
export default class Reviews extends Vue {
  @Prop()
  place = "Kuala Lumpur";
}
</script>

<style scoped>
</style>
