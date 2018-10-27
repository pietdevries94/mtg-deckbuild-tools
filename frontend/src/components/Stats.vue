<template>
  <v-layout row wrap class="mt-3" v-show="entries.length">
    <v-flex xs3>
      <v-card>
        <v-card-title><h4>Types</h4></v-card-title>
        <v-divider />
        <v-list dense>
          <v-list-tile v-for="(amount, key) in stats.types" :key="key">
            <v-list-tile-content>{{key}}</v-list-tile-content>
            <v-list-tile-content class="align-end">{{amount}}</v-list-tile-content>
          </v-list-tile>
          <v-divider />
          <v-list-tile>
            <v-list-tile-content>Total cards:</v-list-tile-content>
            <v-list-tile-content class="align-end">{{entries.length}}</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
    <v-flex xs3>
      <v-card>
        <v-card-title><h4>Tags</h4></v-card-title>
        <v-divider />
        <v-list dense>
          <v-list-tile v-for="(amount, key) in stats.tags" :key="key">
            <v-list-tile-content>{{key}}</v-list-tile-content>
            <v-list-tile-content class="align-end">{{amount}}</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
    <v-flex xs3>
      <v-card>
        <v-card-title><h4>Converted Mana Cost</h4></v-card-title>
        <v-divider />
        <v-list dense>
          <v-list-tile v-for="(amount, key) in stats.cmc" :key="key">
            <v-list-tile-content>{{key}}</v-list-tile-content>
            <v-list-tile-content><v-progress-linear :value="amount*10" /></v-list-tile-content>
            <v-list-tile-content class="align-end">{{amount}}</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
    <v-flex xs3>
      <v-card>
        <v-card-title><h4>Colors</h4></v-card-title>
        <v-divider />
        <v-list dense>
          <v-list-tile v-for="(amount, key) in stats.colorCards" :key="key">
            <v-list-tile-content>{{key}} cards</v-list-tile-content>
            <v-list-tile-content class="align-end">{{amount}}</v-list-tile-content>
          </v-list-tile>
          <v-divider />
          <v-list-tile v-for="(amount, key) in stats.colorSymbols" :key="key">
            <v-list-tile-content>{{key}} casting symbols</v-list-tile-content>
            <v-list-tile-content class="align-end">{{amount}}</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { EntryInterface } from "@/api";

interface statsCat {
  [key: string]: number;
}

interface stats {
  types: statsCat;
  tags: statsCat;
  cmc: statsCat;
  colorCards: statsCat;
  colorSymbols: statsCat;
}

@Component
export default class Stats extends Vue {
  @Prop({ required: true })
  private entries!: EntryInterface[];

  private get stats() {
    const stats: stats = {
      types: {},
      tags: {},
      cmc: {},
      colorCards: {},
      colorSymbols: {}
    };
    let maxCMC = 0;

    this.entries.forEach(entry => {
      // Types
      this.getBaseTypes(entry.card.type_line).forEach(type => {
        stats.types[type] = this.addToUnknown(stats.types[type]);
      });

      // Tags
      entry.tags.forEach(tag => {
        stats.tags[tag] = this.addToUnknown(stats.tags[tag]);
      });

      // CMC
      const cmc = this.getCMC(entry.card.casting_cost);
      stats.cmc[cmc] = this.addToUnknown(stats.cmc[cmc]);

      if (Number(cmc) > maxCMC) maxCMC = Number(cmc);

      // Colors
      const colors = this.getColorIdentities(entry.card.casting_cost);
      const keys = Object.keys(colors);
      keys.forEach(symbol => {
        const name = this.symbolToName(symbol);
        stats.colorCards[name] = this.addToUnknown(stats.colorCards[name]);
        stats.colorSymbols[name] = this.addToUnknown(
          stats.colorSymbols[name],
          colors[symbol]
        );
      });
      if (keys.length === 0) {
        stats.colorCards["Colorless"] = this.addToUnknown(
          stats.colorCards["Colorless"]
        );
      }
    });

    for (let i = 0; i <= maxCMC; i++) {
      if (stats.cmc[i] === undefined) stats.cmc[i] = 0;
    }

    return stats;
  }

  private getBaseTypes(type: string) {
    const types = type.split(" —", 1)[0].split(" ");
    const legendaryI = types.indexOf("Legendary");
    if (legendaryI > -1) {
      types.splice(legendaryI, 1);
    }
    return types;
  }

  private getCMC(castingCost: string) {
    const costArray = this.getManaCostArray(castingCost);
    if (costArray.length === 0) return "No cost";

    let cost = 0;

    costArray.forEach(symbol => {
      if (symbol === "X") return;

      const n = Number(symbol);
      if (!isNaN(n)) {
        cost += n;
      } else {
        cost++;
      }
    });

    return String(cost);
  }

  private getColorIdentities(castingCost: string) {
    const stats: statsCat = {};

    this.getManaCostArray(castingCost).forEach(symbol => {
      symbol.split("/").forEach(symbol => {
        if (!["R", "B", "W", "G", "U"].includes(symbol)) return;

        if (stats[symbol] === undefined) {
          stats[symbol] = 1;
        } else {
          stats[symbol]++;
        }
      });
    });
    return stats;
  }

  private symbolToName(symbol: string) {
    switch (symbol) {
      case "R":
        return "Red";
      case "B":
        return "Black";
      case "W":
        return "White";
      case "G":
        return "Green";
      case "U":
        return "Blue";
      default:
        return "Unknown";
    }
  }

  private getManaCostArray(castingCost: string) {
    return castingCost.slice(1, -1).split("}{");
  }

  private addToUnknown(original: number | undefined, add?: number) {
    if (original === undefined) original = 0;
    if (add === undefined) add = 1;
    return original + add;
  }
}
</script>
