<template>
  <div>
    <select v-model="listID" class="block">
      <option value="0" />
      <option v-for="list in lists" :key="list.id" :value="list.id">
      {{ list.name }} {{ listHasEntry(list.id) ? '&#x2713;' : '' }}
      </option>
    </select>
    New list: <input v-model="newList" /> <span class="clickable" @click="addList()">&#x2795;</span>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { ListInterface, EntryInterface, postList } from "@/add-to-list/api";

@Component
export default class ListPicker extends Vue {
  @Prop({ required: true })
  private entries!: EntryInterface[];
  @Prop({ required: true })
  private lists!: ListInterface[];
  @Prop({ required: true })
  private value!: number;

  private newList = "";

  private get listID() {
    return this.value;
  }
  private set listID(val: number) {
    this.$emit("input", val);
  }

  private listHasEntry(listID: number): boolean {
    for (let i = 0; i < this.entries.length; i++) {
      if (this.entries[i].list_id === listID) {
        return true;
      }
    }
    return false;
  }

  private addList() {
    postList({ name: this.newList })
      .then(res => {
        this.$emit("reloadListsAndTags");
        this.listID = res.id;
        this.newList = "";
      })
      .catch(err => {
        this.$toasted.error("Could not create list");
      });
  }
}
</script>

<style lang="scss" scoped>
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

.clickable {
  cursor: pointer;
}

.block {
  display: block;
  margin-bottom: 8px;
}
</style>
