<template>
  <modal v-model="visible">
    <div class="wrapper">
      <div class="column">
        <card-preview :url="currentCard.thumbnail_url" />
      </div>
      <div class="column">
        <entry-form :card="currentCard" @done="closeModal" />
      </div>
    </div>
  </modal>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from "vue-property-decorator";
import { CardIDs } from "../index";
import { CardInterface, getCardBySetAndNumber } from "../api";
import Modal from "./Modal.vue";
import CardPreview from "./CardPreview.vue";
import EntryForm from "./EntryForm.vue";

@Component({
  components: { Modal, CardPreview, EntryForm }
})
export default class Base extends Vue {
  private visible: boolean = false;
  private currentCard: CardInterface = this.createEmptyCard();

  @Prop({ required: true })
  private currentCardIDs!: CardIDs;

  @Watch("currentCardIDs")
  private onCurrentCardIDsChange(ids: CardIDs) {
    this.getCurrentCardData(ids);
    this.visible = true;
  }

  private getCurrentCardData(ids: CardIDs) {
    this.currentCard = this.createEmptyCard();
    getCardBySetAndNumber(ids.set!, ids.number!).then(card => {
      this.currentCard = card;
    });
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

  private closeModal() {
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
