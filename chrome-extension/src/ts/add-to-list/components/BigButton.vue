<template>
  <button @click="passtrough" :class="[ size, color ]" :disabled="disabled">{{ label }}</button>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component
export default class BigButton extends Vue {
  @Prop({ required: true })
  private label!: string;

  @Prop({ default: false })
  private disabled!: boolean;

  @Prop({ default: "lg" })
  private size!: string;

  @Prop({ default: "green" })
  private color!: string;

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
  padding: 0;
  color: white;
  text-align: center;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.25);
  border: 0;
  cursor: pointer;
  margin-bottom: 8px;

  &:active {
    top: 1px;
    outline: none;
    -webkit-box-shadow: none;
    box-shadow: none;
  }
  &:disabled,
  &.red:disabled,
  &.green:disabled {
    color: #454545;
    background: #ecf0f1;
    border-bottom: 2px solid #dadedf;
    cursor: default;
    -webkit-box-shadow: inset 0 -2px #dadedf;
    box-shadow: inset 0 -2px #dadedf;
  }

  &.lg {
    height: 60px;
    font-size: 22px;
  }

  &.md {
    height: 40px;
    font-size: 18px;
  }

  &.sm {
    height: 25px;
    font-size: 14px;
  }

  &.red {
    background: #d35400;
    border-bottom: 2px solid #c54e00;
    -webkit-box-shadow: inset 0 -2px #c54e00;
    box-shadow: inset 0 -2px #c54e00;
  }

  &.green {
    background: #2ecc71;
    border-bottom: 2px solid #28be68;
    -webkit-box-shadow: inset 0 -2px #28be68;
    box-shadow: inset 0 -2px #28be68;
  }
}
</style>