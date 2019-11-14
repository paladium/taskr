<template>
  <div class="home">
    <div class="container">
      <h1>Board</h1>
      <div class="row">
        <div class="column"></div>
        <div class="column column-offset-75">
          <button class="button">Add task</button>
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
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Task from "@/components/Task.vue";

@Component({
  components: { Task }
})
export default class Home extends Vue {}
</script>
