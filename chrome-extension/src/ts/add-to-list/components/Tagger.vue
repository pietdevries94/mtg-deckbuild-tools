<template>
  <div>
    <div class="block">
      <label>List tags:</label>
      <select v-model="listTagsInput">
        <option value=""></option>
        <option v-for="(val, index) in availableCurrentListTags" :key="index" :value="val">{{ val }}</option>
      </select>
    </div>

    <div class="block">
      <label>All tags:</label>
      <select v-model="allTagsInput">
        <option value=""></option>
        <option v-for="val in availableTags" :key="val" :value="val">{{ val }}</option>
      </select>
    </div>

    <div class="block">
      <label>New tag:</label>
      <input v-model="newTagInput" /> <span class="clickable" @click="addTagFromInput(newTagInput)">&#x2795;</span>
    </div>

    <div class="tag-container">
      <div class="tag" v-for="(tag, index) in model" :key="index" :value="tag">
        {{tag}}<span class="clickable remove" @click="removeTagByIndex(index)">&#x00D7;</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";

@Component
export default class Tagger extends Vue {
  @Prop({ required: true })
  private allTags!: string[];

  @Prop({ required: true })
  private currentListTags!: string[];

  @Prop({ required: true })
  private value!: string[];

  private listTagsInput = "";
  private allTagsInput = "";
  private newTagInput = "";

  private get model(): string[] {
    return this.value;
  }

  private set model(val: string[]) {
    this.$emit("input", val);
  }

  private get availableCurrentListTags(): string[] {
    return this.currentListTags.filter(tag => this.model.indexOf(tag) === -1);
  }

  private get availableTags(): string[] {
    return this.allTags.filter(tag => this.value.indexOf(tag) === -1);
  }

  private removeTagByIndex(index: number) {
    this.model.splice(index);
  }

  @Watch("listTagsInput")
  @Watch("allTagsInput")
  private addTagFromInput(tag: string) {
    if (tag === "") return;
    this.model.push(tag);
    this.listTagsInput = "";
    this.allTagsInput = "";
    this.newTagInput = "";
  }
}
</script>

<style lang="scss" scoped>
div.tag-container {
  border-width: 1px;
  border-style: solid;
  border-color: rgb(166, 166, 166);
  width: 100%;
  min-height: 40px;
  height: auto;
  padding: 2px;

  div.tag {
    margin: 2px;
    display: inline-block;
    padding: 6px;
    background: #d35400;
    color: white;

    .remove {
      margin-left: 4px;
    }
  }
}

.clickable {
  cursor: pointer;
}

.block {
  display: block;
  margin-bottom: 8px;

  label {
    display: inline-block;
    width: 100px;
  }
}

select {
  -webkit-appearance: menulist;
  width: 197px;
}

input {
  -webkit-appearance: textfield;
  height: 23px;
}

input,
select {
  background-color: rgb(248, 248, 248);
  border-width: 1px;
  border-style: solid;
  border-color: rgb(166, 166, 166);
  border-image: initial;
}
</style>

