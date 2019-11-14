import Vue from 'vue';
import Vuex, { StoreOptions } from 'vuex';

Vue.use(Vuex);

export interface TaskrStore {
  addTaskModalOpened: boolean;
}

const store: StoreOptions<TaskrStore> = {
  state: {
    addTaskModalOpened: false
  },
  mutations: {
    setAddTaskModalOpened(state, value: boolean){
      state.addTaskModalOpened = value;
    }
  }
};

export default new Vuex.Store(store);
