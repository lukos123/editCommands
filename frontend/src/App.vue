<script lang="ts" setup>
import { ref, onMounted } from "vue";
import Json from "./components/Json.vue";
import Commands from "./components/Commands.vue";
import { GetJson, SaveJson } from "./../wailsjs/go/main/App";
import { Quit } from "../wailsjs/runtime/runtime";
// import { main } from "../wailsjs/go/models";
import { CommandType } from "./types/types";
const json = ref<CommandType[]>([]);
const jsonTemp = ref<CommandType[]>([]);
const saving = ref<boolean>(false);
const modalSave = ref<boolean>(false);
const saved = ref<boolean>(false);
const reload = () => {
  GetJson().then((i) => {
    json.value = i;
  });
};
const save = () => {
  modalSave.value = true;
  saving.value = true;
  SaveJson(json.value).then((i) => {
    saving.value = false;

    saved.value = i;
    setTimeout(() => {
      modalSave.value = false;
    }, 1000);
  });
};
const buttonFunction: (commandsArr: string | CommandType[]) => void = (
  commandsArr
) => {
  if (typeof commandsArr != "string" && typeof jsonTemp.value != "string") {
    jsonTemp.value?.forEach((command) => {
      commandsArr.push({ ...command });
    });
  }
};
const buttonFunctionRevers: (commandsArr: string | CommandType[]) => void = (
  commandsArr
) => {
  if (typeof commandsArr != "string" && typeof jsonTemp.value != "string") {
    commandsArr.forEach((command) => {
      jsonTemp.value.push({ ...command });
    });
  }
};
onMounted(reload);
</script>

<template>
  <div id="drag">
    <Transition name="modal">
      <div v-show="modalSave" class="modal" @click="modalSave = false">
        <div v-if="!saving">{{ saved ? "saved" : "err" }}</div>
        <div v-else class="load"></div>
      </div>
    </Transition>
    <button class="reload" @click="reload()">reload</button>
    <button class="save" @click="save()">save</button>

    <div @click="Quit()" class="exit">+</div>
  </div>
  <div class="main">
    <div class="left">
      <Json style="height: 50%; overflow-x: hidden" v-model="jsonTemp" />

      <Commands v-model="jsonTemp" />
    </div>
    <div class="right">
      <Json
        style="height: 100%; overflow-x: hidden"
        v-model="json"
        :button-function="buttonFunction"
        :button-function-revers="buttonFunctionRevers"
      />
    </div>
  </div>
</template>

<style>
#drag {
  width: 100%;
  height: 5vh;
  user-select: none;
  display: flex;
  background-color: #a5a5a5;
  justify-content: right;
  align-items: center;
  button {
    margin: 5px;
  }
}
.load {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  border-top: 1px solid #cccaca;
  animation: load 1s ease infinite;
}
.modal {
  padding: 20px;
  transform-origin: left top;

  position: absolute;
  background-color: #404040;
  color: rgb(81, 255, 0);
  font-size: 30px;
  top: 50%;
  left: 50%;
  transform: translateX(-50%) translateY(-50%);
}
@keyframes load {
  0% {
    transform: rotate(-360deg);
  }
  100% {
    transform: rotate(0deg);
  }
}
.reload,
.save {
  padding: 3px;
  border-radius: 3px;
  transition: 0.2s;
  &:hover {
    transform: scale(1.1);
  }
  &:active {
    transform: scale(0.9);
  }
}

.exit {
  color: rgb(255, 0, 0);
  margin: 10px;
  transform: rotateZ(45deg);
  font-size: 30px;
  font-weight: bold;
  transition: 0.2s;
  &:hover {
    transform: rotateZ(-45deg) scale(1.2);
  }
}
.main {
  width: 100%;
  height: 95vh;
  display: flex;
  --wails-draggable: no-drag;
  background-color: #00000067;
  color: white;
}

.left {
  flex: 2;
}
.right {
  flex: 3;
}
</style>
<style>
@keyframes modal {
  0% {
    opacity: 0;
    transform: scale(0) translateX(-50%) translateY(-50%);
  }
  100% {
    opacity: 1;
    transform: scale(1) translateX(-50%) translateY(-50%);
  }
}
.modal-enter-active {
  animation: modal 0.2s;
}
.modal-leave-active {
  animation: modal 0.2s reverse;
}
</style>
