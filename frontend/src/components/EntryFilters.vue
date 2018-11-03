<template>
  <v-layout row wrap>
    <v-flex offset-xs2 xs2>
      <v-select
        v-model="filters.board"
        :items="['main', 'side', 'maybe', '']"
        label="Board"
        prepend-icon="format_list_numbered"
        clearable
        multiple
      />
    </v-flex>
    <v-flex xs2>
      <v-select
        v-model="filters.included"
        :items="list ? list.included_tags : []"
        label="Include"
        prepend-icon="add_circle"
        clearable
        multiple
      />
    </v-flex>
    <v-flex xs2>
      <v-select
        v-model="filters.excluded"
        :items="list ? list.included_tags : []"
        label="Exclude"
        prepend-icon="remove_circle"
        clearable
        multiple
      />
    </v-flex>
    <v-flex xs2>
      <v-select
        v-model="filters.owned"
        :items="['Yes', 'No']"
        label="Owned"
        prepend-icon="folder_open"
        clearable
      />
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

  private filters = this.cleanFilters();

  @Watch("list")
  private clearFilters() {
    this.filters = this.cleanFilters();
  }

  private cleanFilters(): {
    included: string[];
    excluded: string[];
    board: string[];
    owned: string;
  } {
    return {
      included: [],
      excluded: [],
      board: [],
      owned: ""
    };
  }

  @Watch("filters", { deep: true })
  @Watch("entries", { deep: true })
  private filterEntries() {
    const filtered = this.entries.filter(entry => {
      if (
        this.filters.board.length > 0 &&
        !this.filters.board.includes(entry.board)
      )
        return false;
      for (let i = 0; i < this.filters.included.length; i++) {
        if (!entry.tags.includes(this.filters.included[i])) return false;
      }
      for (let i = 0; i < this.filters.excluded.length; i++) {
        if (entry.tags.includes(this.filters.excluded[i])) return false;
      }
      if (this.filters.owned === "Yes" && entry.card.copies_owned === 0)
        return false;
      if (this.filters.owned === "No" && entry.card.copies_owned !== 0)
        return false;

      return true;
    });
    this.$emit("input", filtered);
  }
}
</script>
