<script setup lang="ts">
// import { ref } from "vue";
const visible = defineModel("visible", { default: false });
const content = defineModel("content", { default: "asd" });
const { isHelp } = defineProps<{ isHelp: boolean }>();
const vFocus = {
  mounted: (el: HTMLInputElement) => {
    setTimeout(() => {
      el.focus();
      el.scrollIntoView({ behavior: "smooth" });
    }, 0);
  },
};
</script>

<template>
  <span v-bind="$attrs" v-if="visible">
    {{ isHelp ? "[" : ""
    }}<input
      v-focus
      @focusout="visible = false"
      @keyup.enter="visible = false"
      type="text"
      v-model.lazy="content"
    />{{ isHelp ? "]" : "" }}
  </span>
  <span v-bind="$attrs" v-else @dblclick="visible = true"
    >{{ (isHelp ? "[" : "") + content + (isHelp ? "]" : "") }}
  </span>
</template>

<style scoped></style>
