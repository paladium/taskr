<template>
  <div class="task">
      <p>
        📄 {{task.name}}
      </p>
      <button class="button button-small button-black" v-if="task.deadline">{{task.deadline}}</button>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class Task extends Vue {
  @Prop()
  task!: Task;
}
</script>
<style lang="css" scoped>
.task {
    border: 1px solid #efdcdc;
    border-radius: 5px;
    padding: 10px;
}
</style>