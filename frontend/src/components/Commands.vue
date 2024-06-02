<script setup lang="ts">
import { ref, watchEffect } from "vue";
import { CommandType } from "../types/types";
const json = defineModel<CommandType[]>({ default: [] });
const commands = ref<string>("");
const before = ref<string>("");
const between = ref<string>("");
const after = ref<string>("\n");
const convert = () => {
  let bef = " +";
  let bet = " +";
  let aft = "";
  if (before.value !== "") {
    bef = before.value;
  }
  if (between.value !== "") {
    bet = between.value;
  }
  if (after.value !== "") {
    aft = after.value;
  }
  let regTxtForCommand = `^${bef}(-*[\\w-]+),*${bet}(.+)$`;
  console.log(between.value);
  console.log(regTxtForCommand);

  if (aft) {
    regTxtForCommand = `^${bef}(-*[\\w-]+),*${bet}(.+)${aft}`;
  }
  let regForCommand = new RegExp(regTxtForCommand, "gm");

  // let regForCommand = new RegExp(`^${bef}+(\w+)${aft}+`, "gm");
  // let regForHelp = new RegExp(`^${bef}+\\w+${bet}+(.+)$`, "gm");
  // if (aft !== "") {
  //   regForHelp = new RegExp(`^${bef}+\\w+${bet}+(.+)${aft}`, "gm");
  // }
  const arr: CommandType[] = [];
  [...commands.value.matchAll(regForCommand)].forEach((command) => {
    arr.unshift({
      c: command[1],
      h: command[2],
      awcf: [],
      cf: "",
      n: [],
      na: false,
      f: [],
      o: false,
    });
  });
  json.value = arr.reverse();
  // const commandHelpArr = [...commands.value.matchAll(regForHelp)]
  // commandArr;
  // json.value.unshift();
  // console.log(commandArr);
  // console.log();

  // console.log(`^[${bef}]+(.+)[${bet}]+`);
};
</script>

<template>
  <div class="commands">
    <input type="text" placeholder="before" v-model="before" />
    <input type="text" placeholder="between" v-model="between" />
    <input type="text" placeholder="after" v-model="after" />
    <button @click="convert">convert</button>
    <!-- <input type="text" placeholder="after" v-model="after" /> -->
    <textarea v-model="commands"></textarea>
  </div>
</template>

<style scoped>
.commands {
  height: 50%;
  border: 1px solid #242424b2;
  background-color: #000000b2;
  color: white;
  display: flex;
  flex-wrap: wrap;

  textarea,
  input {
    color: white;
    border: 1px solid #242424b2;
    background-color: #ffffff00;
    padding: 2px;
  }
  textarea {
    height: 90%;
    resize: none;
    width: 100%;
  }
  input {
    height: 10%;
    width: 22.22%;
  }
  button {
    transition: 0.2s;
    height: 10%;
    width: 33.33%;
    background-color: #d6d6d6;
    &:hover {
      background-color: #ffffff;
    }
    &:active {
      background-color: #868686;
    }
  }
}
</style>
