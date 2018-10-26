<template>
  <div>
    <list-picker v-model="currentList" />
    <div v-show="currentList != null">
      <entries-table :entries="entries" :list="currentList" @reload:entries="loadEntriesForList(currentList.id)" />
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

@Component({
  components: { EntriesTable, ListPicker }
})
export default class Home extends Vue {
  private currentList: ListInterface | null = null;
  private entries: EntryInterface[] = [];

  private async loadEntriesForList(listID: number) {
    const res = await getListEntries(listID);
    this.entries = res.entries;
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
