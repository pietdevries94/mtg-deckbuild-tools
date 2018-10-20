<template>
  <div>
    <h2>Card:</h2>
    <div>
      {{ this.card }}
    </div>
    <h2>lists:</h2>
    <div>
      {{ this.lists }}
    </div>
    <h2>tags:</h2>
    <div>
      {{ this.tags }}
    </div>
    <h2>entries:</h2>
    <div>
      {{ this.entries }}
    </div>
    <BigButton @click="addEntry" label="Add to list" :disabled="false" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import {
  CardInterface,
  postEntry,
  ListInterface,
  EntryInterface
} from "@/add-to-list/api";
import BigButton from "./BigButton.vue";

@Component({
  components: { BigButton }
})
export default class EntryForm extends Vue {
  @Prop({ required: true })
  private card!: CardInterface;
  @Prop({ required: true })
  private entries!: EntryInterface;
  @Prop({ required: true })
  private lists!: ListInterface[];
  @Prop({ required: true })
  private tags!: string[];

  private async addEntry() {
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


