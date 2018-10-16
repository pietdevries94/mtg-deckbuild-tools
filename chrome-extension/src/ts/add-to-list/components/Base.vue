<template>
  <modal v-model="visible">
    <div class="column">
      foo
    </div>
    <div class="column">
      {{ currentCard }}
    </div>
  </modal>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from "vue-property-decorator";
import { Card } from "../index";
import Modal from "./Modal.vue";

@Component({
  components: { Modal }
})
export default class Base extends Vue {
  private visible: boolean = false;

  @Prop()
  private currentCardIDs!: Card;

  @Watch("currentCardIDs")
  private onCurrentCardIDsChange() {
    this.visible = true;
  }
}
</script>

<style lang="scss" scoped>
.column {
  float: left;
  margin: 0 15px;
}
</style>

