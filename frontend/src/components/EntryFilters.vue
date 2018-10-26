<template>
  <v-layout row wrap>
    <v-flex xs2 v-for="(tag, index) in list ? list.included_tags : []" :key="index">
      <v-select
        v-model="filters[tag]"
        :items="['Yes', 'No']"
        :label="tag"
        clearable
      ></v-select>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
import { EntryInterface, ListInterface } from "@/api";

@Component
export default class EntryFilters extends Vue {
  @Prop({ required: true })
  private entries!: EntryInterface[];
  @Prop({ required: true })
  private list!: ListInterface;

  private filters: { [key: string]: string } = {};

  @Watch("list")
  private clearFilters() {
    this.filters = {};
  }

  @Watch("filters", { deep: true })
  @Watch("entries", { deep: true })
  private filterEntries() {
    const required: string[] = [];
    const forbidden: string[] = [];

    for (let tag in this.filters) {
      if (this.filters[tag] === "Yes") {
        required.push(tag);
      } else if (this.filters[tag] === "No") {
        forbidden.push(tag);
      }
    }

    const filtered = this.entries.filter(entry => {
      for (let i = 0; i < required.length; i++) {
        if (!entry.tags.includes(required[i])) return false;
      }
      for (let i = 0; i < forbidden.length; i++) {
        if (entry.tags.includes(forbidden[i])) return false;
      }
      return true;
    });
    this.$emit("input", filtered);
  }
}
</script>
