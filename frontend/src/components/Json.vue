<script setup lang="ts">
import { ref, computed } from "vue";
// import { main } from "../../wailsjs/go/models";
import Command from "./Command.vue";
import { CommandType } from "../types/types";

const json = defineModel<CommandType[]>({ default: [] });

const filter = ref<string>("");

const jsonFiltered = computed<CommandType[]>(() => {
  if (filter.value == "") {
    return json.value;
  }
  const temp: CommandType[] = [];
  json.value.forEach((i) => {
    // console.log(i.c, filter.value, i.c.includes(filter.value));

    if (i.c.includes(filter.value)) {
      // console.log(i);

      temp.unshift(i);
    }
  });
  // console.log(temp);

  return temp;
  // console.log(tempJson);
});
const { buttonFunction,buttonFunctionRevers } = defineProps<
{
  buttonFunction?: (commandsArr: CommandType[] | string) => void;
  buttonFunctionRevers?: (commandsArr: CommandType[] | string) => void;
}
>();
// console.log(json);
const add = () => {
  // console.log(json.value);

  json.value.unshift({
    c: "newCommmand",
    h: "",
    awcf: [],
    cf: "",
    n: [],
    na: false,
    o:false,
    f:[]
  });
  // console.log(json.value);
};
</script>

<template>
  <div class="commands">
    <div class="top-panel">
      <button class="button-add" @click="add">+</button>
      <div class="filter"><input type="text" v-model="filter" /></div>
    </div>
    <TransitionGroup name="list">
      <Command
        v-for="(obj, id) in jsonFiltered"
        :new="obj.c == 'newCommmand'"
        :key="obj.c + obj.cf"
        :button-function="buttonFunction"
        :button-function-revers="buttonFunctionRevers"
        v-model="jsonFiltered[id]"
      >
        
        <span @click="json.splice(json.indexOf(jsonFiltered[id]), 1)">del</span>
      </Command>
    </TransitionGroup>
    <!-- <Command :key="json[19 - 5].c" :json="json[19 - 5]" /> -->
  </div>
</template>
<style>
.list-move, /* применять переход к движущимся элементам */
.list-enter-active,
.list-leave-active {
  transition: all 0.2s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}


.list-leave-active {
  position: absolute;
}
@keyframes commands-in {
  0% {
    opacity: 0;
    opacity: 0;
    transform: translateY(30px);
  }
  100% {
    opacity: 1;

    /* transform: translateX(30px); */
  }
}
.commands-in-enter-active {
  animation: commands-in 0.2s;
}
.commands-in-leave-active {
  animation: commands-in 0.2s reverse;
}
</style>
<style scoped>
.commands {
  overflow: auto;

  border: 1px solid #242424b2;
  background-color: #000000b2;
}
.top-panel {
  display: flex;
  justify-content: space-around;
}
.button-add {
  margin-left: 10px;
  width: 20px;
  height: 20px;
  padding: 5px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
  margin-top: 3px;
  margin-bottom: 3px;
}
</style>
