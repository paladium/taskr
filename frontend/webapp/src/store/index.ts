import Vue from 'vue';
import Vuex, { StoreOptions } from 'vuex';
import { Task } from '@/models/task';
import { container } from 'tsyringe';
import TaskService from '@/services/TaskService';

Vue.use(Vuex);

export interface TaskrStore {
  addTaskModalOpened: boolean;
  tasks: Array<Task>;
  newTask: Task | null;
}

const store: StoreOptions<TaskrStore> = {
  state: {
    addTaskModalOpened: false,
    tasks: [],
    newTask: null
  },
  mutations: {
    setAddTaskModalOpened(state, value: boolean) {
      state.addTaskModalOpened = value;
    },
    setTasks(state, tasks: Array<Task>) {
      state.tasks = tasks;
    },
    newTask(state){
      state.newTask = <Task>{
        id: "",
        name: "",
        section: "backlog"
      };
    }
  },
  actions: {
    tasks({commit}) {
      return container.resolve(TaskService).tasks().then(tasks => { commit('setTasks', tasks); });
    },
    moveTask({commit}, task: Task) {
      return container.resolve(TaskService).moveTask(task);
    },
    addTask({commit}, task: Task) {
      return container.resolve(TaskService).addTask(task);
    }
  }
};

export default new Vuex.Store(store);
