<template>
  <modal v-model="visible">
    <div class="wrapper">
      <div class="column">
        <card-preview :url="currentCard.thumbnail_url" />
      </div>
      <div class="column">
        <entry-form :card="currentCard" :lists="lists" :tags="tags" :entries="currentCardEntries" @done="onSubmit" @reloadListsAndTags="loadListsAndTags" />
      </div>
    </div>
  </modal>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from "vue-property-decorator";
import { CardIDs } from "../index";
import {
  CardInterface,
  getCardBySetAndNumber,
  EntryInterface,
  ListInterface,
  getLists,
  getTags,
  getCardByName
} from "../api";
import Modal from "./Modal.vue";
import CardPreview from "./CardPreview.vue";
import EntryForm from "./EntryForm.vue";
import { AxiosError } from "axios";

@Component({
  components: { Modal, CardPreview, EntryForm }
})
export default class Base extends Vue {
  private visible: boolean = false;
  private currentCard: CardInterface = this.createEmptyCard();
  private currentCardEntries: EntryInterface[] = [];
  private lists: ListInterface[] = [];
  private tags: string[] = [];
  private validCache = false;

  @Prop({ required: true })
  private currentCardIDs!: CardIDs;

  @Prop({ required: true })
  private currentCardName!: string;

  @Watch("currentCardIDs")
  private async onCurrentCardIDsChange(ids: CardIDs) {
    if (!this.validCache) {
      this.loadListsAndTags();
    }

    const success = await this.getCurrentCardDataByIDs(ids);
    if (success) {
      this.visible = true;
    }
  }

  @Watch("currentCardName")
  private async onCurrentCardNameChange(name: string) {
    if (!this.validCache) {
      this.loadListsAndTags();
    }

    const success = await this.getCurrentCardDataByName(name);
    if (success) {
      this.visible = true;
    }
  }

  private async loadListsAndTags() {
    try {
      const listRes = await getLists();
      this.lists = listRes.lists;
      const tagRes = await getTags();
      this.tags = tagRes.tags;
    } catch (e) {
      this.axiosErrorToast(e);
    }
  }

  private axiosErrorToast(e: any) {
    const error = <AxiosError>e;
    let msg: string;
    if (error.code == undefined) {
      msg = "Backend is not running!";
    } else {
      msg = "Could not load card info due to an unknown error. :(";
    }

    this.$toasted.error(msg);

    return false;
  }

  private async getCurrentCardDataByIDs(ids: CardIDs) {
    this.currentCard = this.createEmptyCard();
    try {
      const res = await getCardBySetAndNumber(ids.set, ids.number);
      this.currentCard = res.card;
      this.currentCardEntries = res.entries;

      return true;
    } catch (e) {
      this.axiosErrorToast(e);
      return false;
    }
  }

  private async getCurrentCardDataByName(name: string) {
    this.currentCard = this.createEmptyCard();
    try {
      const res = await getCardByName(name);
      this.currentCard = res.card;
      this.currentCardEntries = res.entries;

      return true;
    } catch (e) {
      this.axiosErrorToast(e);
      return false;
    }
  }

  private createEmptyCard(): CardInterface {
    return {
      scryfall_id: "",
      set_code: "",
      set_number: "",
      name: "",
      oracle_id: "",
      thumbnail_url: chrome.runtime.getURL("images/card-placeholder.jpeg"),
      updated_at: ""
    };
  }

  private onSubmit() {
    this.validCache = false;
    this.visible = false;
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
  flex-wrap: wrap;

  .column {
    flex: 1 1 200px;
  }
}
</style>
