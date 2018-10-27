<template>
  <v-layout row wrap>
    <v-flex xs2><div class="subheading mt-3">Add tags</div></v-flex>
    <v-flex xs2>
      <v-select
        v-model="tagSelect"
        :items="unusedTags"
        label="Unused tags"
        solo
      />
    </v-flex>
    <v-flex xs2>
      <v-text-field
        v-model="newTag"
        :append-icon="newTag ? 'add_circle' : ''"
        @click:append="onAddClick"
        @keyup.enter="onAddClick"
        @click:clear="newTag = null"
        label="New tag"
        clearable
        solo
      />
    </v-flex>

  </v-layout>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
import { getTags } from "@/api";

@Component
export default class ListTagger extends Vue {
  @Prop({ required: true })
  private value!: string[];
  private allTags: string[] = [];

  private tagSelect = "";
  private newTag = "";

  @Watch("tagSelect")
  private addTag(tag: string) {
    if (tag === "") return;

    this.currentTags.push(tag);
    this.tagSelect = "";
    this.newTag = "";
  }

  private onAddClick() {
    this.addTag(this.newTag);
  }

  private get currentTags() {
    return this.value;
  }
  private set currentTags(val) {
    this.$emit("input", val);
  }

  private get unusedTags() {
    return this.allTags.filter(tag => !this.currentTags.includes(tag));
  }

  private async loadAllTags() {
    const res = await getTags();
    this.allTags = res.tags;
  }

  private mounted() {
    this.loadAllTags();
  }
}
</script>
