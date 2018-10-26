<template>
  <div>
    <list-picker v-model="currentList" />
    <div v-show="currentList != null">
      <entries-table />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import EntriesTable from "@/components/EntriesTable.vue";
import { getLists, ListInterface } from "@/api.ts";
import ListPicker from "@/components/ListPicker.vue";

@Component({
  components: { EntriesTable, ListPicker }
})
export default class Home extends Vue {
  private currentList: ListInterface | null = null;
}
</script>
