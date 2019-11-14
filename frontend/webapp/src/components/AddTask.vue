<template>
  <div class="add-task modal" v-if="modalOpened">
    <div class="content" v-click-outside="closeModal">
      <input type="text" placeholder="*Name" v-model="task.name" />
      <input type="text" placeholder="Deadline" v-model="task.deadline" />
      <button class="button button-outline" @click="addTask">💾</button>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { TaskrStore } from "../store";

@Component
export default class AddTask extends Vue {
  get modalOpened() {
    return (this.$store.state as TaskrStore).addTaskModalOpened;
  }
  closeModal() {
    this.$store.commit("setAddTaskModalOpened", false);
  }
  get task() {
    return (this.$store.state as TaskrStore).newTask;
  }
  addTask() {
    this.closeModal();
    this.$store.dispatch("addTask", this.task);
  }
}
</script>