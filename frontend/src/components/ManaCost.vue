<template>
  <div>
    <img 
        v-for="(cost, index) in costCircles"
        :src="cost" 
        :key="index"
        class="mana-cost"
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class ManaCost extends Vue {
  @Prop({ required: true })
  private value!: string;

  private get costCircles() {
    if (this.value === "") return [];
    return this.value
      .replace(new RegExp("/", "g"), "")
      .slice(1, -1)
      .split("}{")
      .map(cost => require(`../assets/mana-symbols/${cost}.svg`));
  }
}
</script>

<style scoped>
.mana-cost {
  margin-right: 4px;
  height: 18px;
}
</style>
