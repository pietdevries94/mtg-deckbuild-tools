<template>
  <div>
    <div class="form-group">
      <label>List:</label>
      <list-picker v-model="entry.list_id" :lists="lists" :entries="entries" @reloadListsAndTags="$emit('reloadListsAndTags')" />
    </div>
    <div class="form-group">
      <label>Tags:</label>
      <tagger v-model="entry.tags" :all-tags="tags" :current-list-tags="currentListTags" />
    </div>
    <BigButton @click="addEntry" label="Save" :disabled="!valid" color="green" size="md" />
    <BigButton @click="deleteEntry" label="Delete entry" :disabled="!currentEntryExists" color="red" size="sm" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
import {
  CardInterface,
  postEntry,
  ListInterface,
  EntryInterface,
  deleteEntry
} from "@/add-to-list/api";
import BigButton from "./BigButton.vue";
import Tagger from "./Tagger.vue";
import ListPicker from "./ListPicker.vue";

@Component({
  components: { BigButton, Tagger, ListPicker }
})
export default class EntryForm extends Vue {
  @Prop({ required: true })
  private card!: CardInterface;
  @Prop({ required: true })
  private entries!: EntryInterface[];
  @Prop({ required: true })
  private lists!: ListInterface[];
  @Prop({ required: true })
  private tags!: string[];

  private entry: EntryInterface = { scryfall_id: "", list_id: 0, tags: [] };

  private get currentListTags(): string[] {
    for (let i = 0; i < this.lists.length; i++) {
      if (this.lists[i].id === this.entry.list_id) {
        return this.lists[i].included_tags;
      }
    }
    return [];
  }

  private get valid(): boolean {
    return this.entry.scryfall_id !== "" && this.entry.list_id > 0;
  }

  private get currentEntryExists(): boolean {
    for (let i = 0; i < this.entries.length; i++) {
      if (this.entries[i].list_id === this.entry.list_id) {
        return true;
      }
    }
    return false;
  }

  @Watch("card")
  private prepareEntry(card: CardInterface) {
    this.entry = {
      scryfall_id: card.scryfall_id,
      list_id: 0,
      tags: []
    };

    chrome.storage.local.get(
      ["last_list_id"],
      (res: { last_list_id?: number }) => {
        if (res.last_list_id !== undefined) {
          this.entry.list_id = res.last_list_id;
        }
      }
    );

    this.setExistingTags();
  }

  @Watch("entry.list_id")
  private setExistingTags() {
    for (let i = 0; i < this.entries.length; i++) {
      if (this.entries[i].list_id === this.entry.list_id) {
        this.entry.tags = this.entries[i].tags;
        return;
      }
    }
  }

  private async addEntry() {
    let success: boolean;
    try {
      await postEntry(this.entry);
      chrome.storage.local.set({ last_list_id: this.entry.list_id });
      this.$toasted.success("Entry succesfully added!");
    } catch (e) {
      this.$toasted.error("Something went wrong :(");
    }

    this.$emit("done");
  }

  private async deleteEntry() {
    let success: boolean;
    try {
      await deleteEntry(this.entry.scryfall_id, this.entry.list_id);
      this.$toasted.success("Entry succesfully deleted!");
    } catch (e) {
      this.$toasted.error("Something went wrong :(");
    }

    this.$emit("done");
  }
}
</script>

<style lang="scss" scoped>
.form-group {
  margin-bottom: 16px;

  label {
    display: block;
    margin-bottom: 4px;
    font-weight: 500;
  }
}
</style>

