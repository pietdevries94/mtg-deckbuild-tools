<template>
  <button @click="passtrough" :disabled="disabled">{{ label }}</button>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class BigButton extends Vue {
  @Prop({ required: true })
  private label!: string;

  @Prop({ default: false })
  private disabled!: boolean;

  private passtrough() {
    this.$emit("click");
  }
}
</script>

<style lang="scss" scoped>
/* #Silver Flat Button (http://cssdeck.com/labs/beautiful-flat-buttons)
==================================================*/
button {
  position: relative;
  vertical-align: top;
  width: 100%;
  height: 60px;
  padding: 0;
  font-size: 22px;
  color: white;
  text-align: center;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  background: #d35400;
  border: 0;
  border-bottom: 2px solid #c54e00;
  cursor: pointer;
  -webkit-box-shadow: inset 0 -2px #c54e00;
  box-shadow: inset 0 -2px #c54e00;
  &:active {
    top: 1px;
    outline: none;
    -webkit-box-shadow: none;
    box-shadow: none;
  }
  &:disabled {
    color: #454545;
    background: #ecf0f1;
    border-bottom: 2px solid #dadedf;
    cursor: default;
    -webkit-box-shadow: inset 0 -2px #dadedf;
    box-shadow: inset 0 -2px #dadedf;
  }
}
</style>