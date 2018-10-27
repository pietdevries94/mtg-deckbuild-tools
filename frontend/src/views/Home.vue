<template>
  <div>
    <list-picker v-model="currentList" />
    <div v-show="currentList != null">
      <list-tagger v-model="currentListTags" />
      <entry-filters v-model="filteredEntries" :entries="entries" :list="currentList" />
      <entries-table :entries="filteredEntries" :list="currentList" @reload:entries="loadEntriesForList(currentList.id)" />
      <stats label="List stats" :entries="filteredEntries" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import EntriesTable from "@/components/EntriesTable.vue";
import {
  getLists,
  ListInterface,
  EntryInterface,
  getListEntries
} from "@/api.ts";
import ListPicker from "@/components/ListPicker.vue";
import EntryFilters from "@/components/EntryFilters.vue";
import ListTagger from "@/components/ListTagger.vue";
import Stats from "@/components/Stats.vue";

@Component({
  components: { EntriesTable, ListPicker, EntryFilters, ListTagger, Stats }
})
export default class Home extends Vue {
  private currentList: ListInterface | null = null;
  private entries: EntryInterface[] = [];
  private filteredEntries: EntryInterface[] = [];

  private async loadEntriesForList(listID: number) {
    const res = await getListEntries(listID);
    this.entries = res.entries;
  }

  private get currentListTags() {
    return this.currentList ? this.currentList.included_tags : [];
  }
  private set currentListTags(val) {
    if (this.currentList) this.currentList.included_tags = val;
  }

  @Watch("currentList")
  private onCurrentListChange(
    list: ListInterface | null,
    oldList: ListInterface | null
  ) {
    if (list === null || list === undefined || list === oldList) return;

    this.loadEntriesForList(list.id);
  }
}
</script>
