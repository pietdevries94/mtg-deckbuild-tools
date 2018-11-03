<template>
  <v-card>
    <v-card-title><h4>Export current selection</h4></v-card-title>
    <v-divider />
    <v-btn
        color="brown"
        dark
        block
        @click="copyNames()"
    >
        Copy names
    </v-btn>
    <v-snackbar
        v-model="showSnackbar"
    >
        List copied to clipboard
        <v-btn
            flat
            @click="showSnackbar = false"
        >
            Close
        </v-btn>
    </v-snackbar>
  </v-card>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { EntryInterface } from "@/api";

@Component
export default class CopyCurrentList extends Vue {
  @Prop({ required: true })
  private entries!: EntryInterface[];

  private showSnackbar = false;

  private copyNames() {
    const list = this.entries.map(entry => entry.card.name).join("\n");
    this.copyStringToClipboard(list);
  }

  // https://techoverflow.net/2018/03/30/copying-strings-to-the-clipboard-using-pure-javascript/
  private copyStringToClipboard(str: string) {
    // Create new element
    var el = document.createElement("textarea");
    // Set value (string to be copied)
    el.value = str;
    // Set non-editable to avoid focus and move outside of view
    el.setAttribute("readonly", "");
    el.style.position = "absolute";
    el.style.left = "-9999px";
    document.body.appendChild(el);
    // Select text inside element
    el.select();
    // Copy text to clipboard
    document.execCommand("copy");
    // Remove temporary element
    document.body.removeChild(el);

    this.showSnackbar = true;
  }
}
</script>
