<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="entries"
      hide-actions
    >
      <template slot="items" slot-scope="props">
          <td class="text-xs-left">
            <v-tooltip bottom>
              <span slot="activator">{{ props.item.card.name }}</span>
              <img :src="props.item.card.thumbnail_url" class="preview">
            </v-tooltip>
          </td>

          <td class="text-xs-left" v-for="(tag, index) in list.included_tags" :key="index">
            <v-checkbox :inputValue="props.item.tags.includes(tag)" @change="(state) => {updateTags(props.item, tag, state)}" />
          </td>
          
          <td>
            <v-btn color="primary" @click="openCardScryfall(props.item.card)">Open on Scryfall</v-btn>
            <v-btn color="error" @click="deleteEntry(props.item.scryfall_id, props.item.list_id)">Delete entry</v-btn>
          </td>
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
import {
  EntryInterface,
  CardInterface,
  ListInterface,
  deleteEntry,
  postEntry
} from "@/api";

@Component
export default class EntriesTable extends Vue {
  @Prop({ required: true })
  private list!: ListInterface | null;
  @Prop({ required: true })
  private entries!: EntryInterface[];

  private headers: { text: string; value: string; align?: string }[] = [];

  @Watch("entries")
  private setHeaders() {
    if (!this.list) return;

    this.headers = [{ text: "Name", value: "card.name" }];

    this.list.included_tags.forEach((tag, index) => {
      this.headers.push({
        text: tag,
        value: "card.tags." + index
      });
    });

    this.headers.push({ text: "Actions", value: "id", align: "center" });
  }

  private updateTags(entry: EntryInterface, tag: string, state: boolean) {
    if (state) {
      entry.tags.push(tag);
    } else {
      const index = entry.tags.indexOf(tag);
      entry.tags.splice(index, 1);
    }

    postEntry(entry);
  }

  private openCardScryfall(card: CardInterface) {
    window.open(
      "https://scryfall.com/card/" + card.set_code + "/" + card.set_number,
      "_blank"
    );
  }

  private async deleteEntry(scryfallID: string, listID: number) {
    await deleteEntry(scryfallID, listID);
    this.$emit("reload:entries");
  }

  private created() {
    this.setHeaders();
  }
}
</script>

<style scoped>
.preview {
  width: 400px;
}
</style>
