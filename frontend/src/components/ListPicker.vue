<template>
  <div>
    <v-layout row wrap>
      <v-flex xs5>
        <v-select
          v-model="currentList"
          :items="listSelectOptions"
          label="Pick a list"
          @click:clear="currentList = null"
          clearable
          solo
        />
      </v-flex>
      <v-flex xs2>
        <v-btn 
          color="error"
          v-show="currentList"
          @click="deleteDialog = true"
        >
          Delete current list
        </v-btn>
      </v-flex>
      <v-flex xs5>
        <v-text-field
          v-model="newList"
          :append-icon="newList ? 'add_circle' : ''"
          @click:append="addList"
          @click:clear="newList = null"
          label="New list"
          placeholder="New list"
          clearable
          solo
        />
      </v-flex>
    </v-layout>
    <v-dialog
        v-model="deleteDialog"
        width="500"
      >
      <v-card>
          <v-card-title
            class="headline grey lighten-2"
            primary-title
          >
            Are you sure you want to delete this list?
          </v-card-title>
  
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green"
              flat
              @click="deleteDialog = false"
            >
              No
            </v-btn>
            <v-btn
              color="red"
              flat
              @click="deleteList"
            >
              Yes
            </v-btn>
          </v-card-actions>
        </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { ListInterface, getLists, postList, deleteList } from "@/api";

@Component
export default class ListPicker extends Vue {
  private lists: ListInterface[] = [];
  private newList = "";
  private deleteDialog = false;

  @Prop({ required: true })
  private value!: ListInterface | null;

  private get currentList() {
    return this.value;
  }
  private set currentList(val) {
    this.$emit("input", val);
  }

  private get listSelectOptions() {
    return this.lists.map(list => ({
      text: list.name,
      value: list
    }));
  }

  private async loadLists() {
    const res = await getLists();
    this.lists = res.lists;
    return;
  }

  private getListByID(id: number) {
    return this.lists.find(list => list.id === id);
  }

  private async addList() {
    const res = await postList({ name: this.newList });
    await this.loadLists();
    const list = this.getListByID(res.id);
    this.currentList = list!;
    this.newList = "";
  }

  private async deleteList() {
    await deleteList(this.currentList!.id);
    this.deleteDialog = false;
    this.currentList = null;
  }

  private mounted() {
    this.loadLists();
  }
}
</script>
