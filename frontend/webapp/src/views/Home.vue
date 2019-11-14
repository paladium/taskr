<template>
  <div class="home">
    <div class="container">
      <h1>Board</h1>
      <div class="row">
        <div class="column"></div>
        <div class="column column-offset-75">
          <button class="button" @click="openAddTaskModal()">Add task</button>
        </div>
      </div>
      <div class="row">
        <div class="column column-25">
          <h4>Backlog</h4>
          <task
            :task="{'name': 'cool', 'id': 'nice', 'section': 'backlog', 'deadline': 'tomorrow'}"
          ></task>
          <task
            :task="{'name': 'cool', 'id': 'nice', 'section': 'backlog', 'deadline': 'tomorrow'}"
          ></task>
        </div>
        <div class="column column-25">
          <h4>Selected for development</h4>
        </div>
        <div class="column column-25">
          <h4>In progress</h4>
        </div>
        <div class="column column-25">
          <h4>Done</h4>
        </div>
      </div>
    </div>
    <add-task></add-task>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Task from "@/components/Task.vue";
import AddTask from "@/components/AddTask.vue";
import { TaskrStore } from "../store";
import { Task as TaskModel } from "../models/task";
import {Action} from 'vuex-class'

@Component({
  components: { Task, AddTask }
})
export default class Home extends Vue {

  @Action('tasks')
  loadTasks!: () => Promise<any>;

  openAddTaskModal() {
    this.$store.commit("setAddTaskModalOpened", true);
    this.$store.commit('newTask');
  }

  get tasks(): Array<TaskModel> {
    return (this.$store.state as TaskrStore).tasks;
  }

  get backlog() {
    return this.tasks.filter(task => task.section == "backlog");
  }
  get dev() {
    return this.tasks.filter(task => task.section == "dev");
  }
  get progress() {
    return this.tasks.filter(task => task.section == "progress");
  }
  get done() {
    return this.tasks.filter(task => task.section == "done");
  }
  moveTask(task: Task)
  {
    this.$store.dispatch('moveTask', task);
  }
}
</script>
