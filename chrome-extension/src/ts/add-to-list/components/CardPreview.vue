<template>
  <img :src="url" />
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class CardPreview extends Vue {
  @Prop({ required: true })
  private url!: string;
}
</script>

<style lang="scss" scoped>
img {
  width: 244px;
  height: 340px;
}
</style>

