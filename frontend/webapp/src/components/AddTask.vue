<template>
  <div class="add-task modal" v-if="modalOpened">
    <div class="content" v-click-outside="closeModal">
      <input type="text" placeholder="*Name" />
      <input type="text" placeholder="Deadline"/>
      <button class="button button-outline">💾</button>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { TaskrStore } from "../store";

@Component
export default class AddTask extends Vue {
    get modalOpened(){
        return (this.$store.state as TaskrStore).addTaskModalOpened;
    }
    closeModal(event: any){
        this.$store.commit('setAddTaskModalOpened', false);
    }
}
</script>