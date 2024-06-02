<script setup lang="ts">
// import { main } from "../../wailsjs/go/models";

import { onMounted, ref } from "vue";
import { CommandType } from "../types/types";
import EditInput from "./EditInput.vue";

// const {str} = defineProps<{str:string}>();
const json = defineModel<CommandType>({ default: {} });
const props = defineProps<{
  new: boolean;
  buttonFunction?: (commandsArr: CommandType[] | string) => void;
  buttonFunctionRevers?: (commandsArr: CommandType[] | string) => void;
}>();

// const isVisibly = json.value.o;
const commandInputVisibly = ref(false);
const commandFunctionInputVisibly = ref(false);
const commandFunctionNInputVisibly = ref(false);

const helpInputVisibly = ref(false);
const add = () => {
  !json.value.o ? (json.value.o = true) : "";
  if (json.value.cf !== "") {
    json.value.awcf.unshift({
      c: "newCommmand",
      h: "",
      awcf: [],
      cf: "",
      n: [],
      na: false,
      o: false,
      f: [],
    });
  } else if (Array.isArray(json.value.n)) {
    json.value.n.unshift({
      c: "newCommmand",
      h: "",
      awcf: [],
      cf: "",
      n: [],
      na: false,
      o: false,
      f: [],
    });
  }
};
const addF = () => {
  !json.value.o ? (json.value.o = true) : "";
  if (json.value.n.length > 0) {
    if (Array.isArray(json.value.n)) {
      json.value.awcf = [...json.value.n];
      json.value.n = [];
    }
  }
  json.value.cf = "newCommmand";
};
const addFl = () => {
  json.value.f.unshift({
    c: "newCommmand",
    h: "",
    awcf: [],
    cf: "",
    n: [],
    na: false,
    o: false,
    f: [],
  });
};
onMounted(() => {
  if (props.new) {
    commandInputVisibly.value = true;
  }
});
// const emit = defineEmits([""]);
</script>

<template>
  <div class="command">
    <div class="command-in" @click="json.o = !json.o">
      <span>
        <!-- {{ json.c }} -->
        <EditInput
          :is-help="false"
          class="command-c"
          v-model:content="json.c"
          v-model:visible="commandInputVisibly"
        />
        {{ " " }}
        <!-- {{ json.h }} -->
        <EditInput
          :is-help="true"
          class="command-h"
          v-model:content="json.h"
          v-model:visible="helpInputVisibly"
        />
        {{ " " }}
        {{ " " }}
      </span>

      <span class="slot">
        <span v-if="json.n.length > 0 || json.cf.length > 0">{{
          !json.o ? "↓" : "↑"
        }}</span>

        <span @click.stop="add">+</span>
        <span @click.stop="addF">+Fu</span>
        <span @click.stop="addFl">+Fl</span>
        <span @click.stop="json.na = !json.na">{{
          json.na ? "anT" : "anF"
        }}</span>
        <span
          v-if="
            props.buttonFunctionRevers &&
            typeof json.n != 'string' &&
            json.cf == ''
          "
          @click.stop="
            props.buttonFunctionRevers(typeof json.n != 'string' ? json.n : []);
            !json.o ? (json.o = true) : '';
          "
          >{{ "<" }}M</span
        >
        <span
          v-else-if="props.buttonFunctionRevers"
          @click.stop="
            props.buttonFunctionRevers(json.awcf);
            !json.o ? (json.o = true) : '';
          "
          >{{ "<" }}M</span
        >
        <span
          v-if="
            props.buttonFunction && typeof json.n != 'string' && json.cf == ''
          "
          @click.stop="
            props.buttonFunction(typeof json.n != 'string' ? json.n : []);
            !json.o ? (json.o = true) : '';
          "
          >M</span
        >
        <span
          v-else-if="props.buttonFunction"
          @click.stop="
            props.buttonFunction(json.awcf);
            !json.o ? (json.o = true) : '';
          "
          >M</span
        >

        <slot />
      </span>
    </div>
    <Transition name="commands-in">
      <div class="commands-in" v-if="json.o">
        <TransitionGroup name="list">
          <template v-if="json.cf != ''">
            <template v-if="typeof json.n === 'string'">
              <div style="padding-left: 10px">
                command :
                <EditInput
                  :is-help="false"
                  class="command-c"
                  v-model:content="json.cf"
                  v-model:visible="commandFunctionInputVisibly"
                />
                <span class="del" @click="json.cf = ''"> del</span>
              </div>
              <div style="padding-left: 20px">
                command :<EditInput
                  :is-help="false"
                  class="command-c"
                  v-model:content="json.n"
                  v-model:visible="commandFunctionNInputVisibly"
                />
                <span @click="json.n = []" class="del"> del</span>
              </div>
            </template>
            <template v-else>
              <template v-if="json.na">
                <div style="padding-left: 10px">
                  command :<EditInput
                    :is-help="false"
                    class="command-c"
                    v-model:content="json.cf"
                    v-model:visible="commandFunctionInputVisibly"
                  />
                  <span
                    @click="
                      () => {
                        json.cf = '';
                        json.n = [...json.awcf];
                        json.awcf = [];
                      }
                    "
                    class="del"
                  >
                    del</span
                  >
                </div>

                <Command
                  :button-function-revers="buttonFunctionRevers"
                  :button-function="props.buttonFunction"
                  :new="obj.c == 'newCommmand'"
                  v-for="(obj, id) in json.awcf.map((i) => {
                    i.n = json.n;
                    return i;
                  })"
                  :key="obj.c + obj.h"
                  v-model="
                    json.awcf.map((i) => {
                      i.n = json.n;
                      return i;
                    })[id]
                  "
                >
                  <span
                    @click="
                      json.n.splice(
                        json.n.indexOf(
                          json.awcf.map((i) => {
                            i.n = json.n;
                            return i;
                          })[id]
                        ),
                        1
                      )
                    "
                    >del</span
                  >
                </Command>
              </template>
              <template v-else>
                <div style="padding-left: 10px">
                  command :<EditInput
                    :is-help="false"
                    class="command-c"
                    v-model:content="json.cf"
                    v-model:visible="commandFunctionInputVisibly"
                  />
                  <span
                    @click="
                      () => {
                        json.cf = '';
                        json.n = [...json.awcf];
                        json.awcf = [];
                      }
                    "
                    class="del"
                  >
                    del</span
                  >
                </div>

                <Command
                  :button-function-revers="buttonFunctionRevers"
                  :button-function="props.buttonFunction"
                  :new="obj.c == 'newCommmand'"
                  v-for="(obj, id) in json.awcf"
                  :key="obj.c + obj.h"
                  v-model="json.awcf[id]"
                >
                  <span
                    @click="
                      json.awcf.splice(json.awcf.indexOf(json.awcf[id]), 1)
                    "
                  >
                    del
                  </span>
                </Command>
              </template>
            </template>
          </template>
          <template v-else>
            <template v-if="typeof json.n === 'string'"> </template>
            <template v-else>
              <Command
                :button-function-revers="buttonFunctionRevers"
                :button-function="props.buttonFunction"
                v-for="(obj, id) in json.n"
                :key="obj.c + obj.h"
                v-model="json.n[id]"
                :new="obj.c == 'newCommmand'"
              >
                <span @click="json.n.splice(json.n.indexOf(json.n[id]), 1)"
                  >del</span
                >
              </Command>
            </template>
          </template></TransitionGroup
        >
      </div>
    </Transition>
    <Transition name="commands-in">
      <div class="commands-in flags" v-if="json.o && json.f.length > 0">
        <TransitionGroup name="list">
          <Command
            :button-function-revers="buttonFunctionRevers"
            :button-function="props.buttonFunction"
            v-for="(obj, id) in json.f"
            :key="obj.c + obj.h"
            v-model="json.f[id]"
            :new="obj.c == 'newCommmand'"
          >
            <span @click="json.f.splice(json.f.indexOf(json.f[id]), 1)"
              >del</span
            >
          </Command>
        </TransitionGroup>
      </div>
    </Transition>
  </div>
</template>

<style>
.command {
  /* overflow: auto; */
  user-select: none;

  padding-left: 10px;
}
/* .slot {
  align-self: right;
} */
.commands-in {
  position: relative;
}
.command-c,
.command-h {
  input {
    width: 50px;
    padding: 0;
    background-color: #ffffff00;
    /* color: rgb(0, 102, 255); */
  }
}
.command-c {
  color: rgb(0, 102, 255);
  input {
    color: rgb(0, 102, 255);
  }
}
.command-h {
  color: rgb(212, 255, 0);
  width: 200px;
  /* скрывает текст, который не помещается */

  input {
    color: rgb(212, 255, 0);
    width: 200px;
  }
}
.commands-in::before {
  content: "";
  background-color: #b4b4b4;
  position: absolute;
  height: 100%;
  width: 1px;
  top: 0;
  left: 0;
  display: block;
}
.flags::after {
  content: "";
  background-color: #b4b4b4;
  position: absolute;
  width: 100%;
  height: 1px;
  top: 0;
  left: 0;
  display: block;
}
.command-in {
  transition: 0.2s;
  padding: 2px;
  width: 100%;
  display: flex;
  justify-content: space-between;
  & > span {
    display: inline-block;
  }
  * {
    transition: 0.2s;
  }
  & > .slot {
    /* width: 80%; */
    text-align: right;
    align-self: center;
    white-space: nowrap;

    :hover {
      background-color: #fdfdfd;
      color: black;
    }
    span {
      display: inline-block;
      border-radius: 5px;
      padding-right: 3px;
      padding-left: 3px;

      /* text-align: center; */
    }
  }
}
.del {
  &:hover {
    background-color: #fdfdfd;
    color: black;
  }

  padding-right: 5px;
  padding-left: 5px;
  /* text-align: center; */
}

.command-in:hover {
  background-color: #b1b1b1;
  .command-c,
  .command-h {
    color: black;
    input {
      color: black;
    }
  }
  & > .slot {
    span {
      color: black;
    }
  }
  .command-h {
    background-color: rgb(212, 255, 0);
  }
  .command-c {
    background-color: rgb(0, 102, 255);
  }
}
</style>
