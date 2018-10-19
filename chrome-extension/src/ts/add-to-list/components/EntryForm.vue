<template>
  <div>
    {{ this.card }}
    <BigButton @click="addEntry" label="Add to list" :disabled="false" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { CardInterface, postEntry } from "@/add-to-list/api";
import BigButton from "./BigButton.vue";

@Component({
  components: { BigButton }
})
export default class EntryForm extends Vue {
  @Prop({ required: true })
  private card!: CardInterface;

  private async addEntry() {
    console.log(this.card.scryfall_id);
    let success: boolean;
    try {
      await postEntry({
        scryfall_id: this.card.scryfall_id
      });
      this.$toasted.success("Entry succesfully added!");
    } catch (e) {
      this.$toasted.error("Something went wrong :(");
    }

    this.$emit("done");
  }
}
</script>


