<template>
  <div>
    <v-data-table
      :headers="headers"
      :items="entries"
      hide-actions
    >
      <template slot="items" slot-scope="props">
          <td class="text-xs-left">
            <v-tooltip bottom>
              <span slot="activator">{{ props.item.card.name }}</span>
              <img :src="props.item.card.thumbnail_url" class="preview">
            </v-tooltip>
          </td>

          <td class="text-xs-left"><mana-cost :value="props.item.card.casting_cost" /></td>
          <td class="text-xs-left">{{ props.item.card.type_line }}</td>

          <td class="text-xs-left" v-for="(tag, index) in list.included_tags ? list.included_tags : []" :key="index">
            <v-checkbox
              class="checkbox"
              :inputValue="props.item.tags.includes(tag)"
              @change="(state) => {updateTags(props.item, tag, state)}"
            />
          </td>

          <td class="text-xs-left">€ {{ props.item.card.online_price }}</td>
          
          <td>
            <v-btn
              color="purple"
              small
              dark
              @click="openCardScryfall(props.item.card)"
            >
              Scryfall
            </v-btn>
            <v-btn
              color="blue"
              small
              dark
              @click="openLink(props.item.card.buy_link)"
            >
              CardMarket
            </v-btn>
            <v-btn
              class="deletebtn"
              color="error"
              small
              @click="deleteEntry(props.item.scryfall_id, props.item.list_id)"
            >
              Delete
            </v-btn>
          </td>
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
import {
  EntryInterface,
  CardInterface,
  ListInterface,
  deleteEntry,
  postEntry
} from "@/api";
import ManaCost from "@/components/ManaCost.vue";

@Component({
  components: { ManaCost }
})
export default class EntriesTable extends Vue {
  @Prop({ required: true })
  private list!: ListInterface | null;
  @Prop({ required: true })
  private entries!: EntryInterface[];

  private headers: {
    text: string;
    value: string;
    align?: string;
    sortable?: boolean;
  }[] = [];

  @Watch("list.included_tags")
  @Watch("entries")
  private setHeaders() {
    if (!this.list) return;

    this.headers = [
      { text: "Name", value: "card.name" },
      { text: "Cost", value: "card.casting_cost" },
      { text: "Type", value: "card.type_line" }
    ];

    this.list.included_tags.forEach((tag, index) => {
      this.headers.push({
        text: tag,
        value: "card.tags." + index,
        sortable: false
      });
    });

    this.headers.push(
      { text: "Price", value: "card.online_price" },
      { text: "Actions", value: "id", align: "center" }
    );
  }

  private updateTags(entry: EntryInterface, tag: string, state: boolean) {
    if (state) {
      entry.tags.push(tag);
    } else {
      const index = entry.tags.indexOf(tag);
      entry.tags.splice(index, 1);
    }

    postEntry(entry);
  }

  private openLink(url: string) {
    window.open(url, "_blank");
  }

  private openCardScryfall(card: CardInterface) {
    this.openLink(
      "https://scryfall.com/card/" + card.set_code + "/" + card.set_number
    );
  }

  private async deleteEntry(scryfallID: string, listID: number) {
    await deleteEntry(scryfallID, listID);
    this.$emit("reload:entries");
  }

  private created() {
    this.setHeaders();
  }
}
</script>

<style scoped>
.preview {
  width: 350px;
  border-radius: 4.75% / 3.5%;
}

.checkbox {
  padding-top: 10px;
  margin-bottom: -10px;
}

.deletebtn {
  margin-left: 20px;
}
</style>
