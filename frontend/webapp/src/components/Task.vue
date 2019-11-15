<template>
  <div class="task">
    <p>📄 {{task.name}}</p>
    <div class="row">
      <div class="column-50">
        <button class="button button-small button-outline" v-if="task.deadline">{{task.deadline}}</button>
      </div>
      <div class="column">
        <button class="button button-small button-black" v-if="task.section != 'done'" @click="moveTask">Mv</button>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";
import { TaskModel } from "../models/task";

@Component
export default class Task extends Vue {
  @Prop()
  task!: TaskModel;

  moveTask(){
    this.$emit('moved', this.task);
  }
}
</script>
<style lang="css" scoped>
.task {
  border: 1px solid #efdcdc;
  border-radius: 5px;
  padding: 10px;
  margin-top: 10px;
}
.row{
  padding-left: 5px;
}
</style>