import {injectable} from 'tsyringe'
import BaseService from './BaseService';
import { TaskModel } from '@/models/task';
import axios from 'axios'
import { BASE, api } from './api';

@injectable()
export default class TaskService extends BaseService
{
    tasks(): Promise<Array<TaskModel>>
    {
        return axios.get(this.joinUrl(BASE, api.tasks)).then(response => response.data.tasks);
    }
    moveTask(task: TaskModel): Promise<any>
    {
        return axios.put(this.joinUrl(BASE, api.tasks, api.moveTask), {id: task.id}).then(response => response.data);
    }
    addTask(task: TaskModel): Promise<any>
    {
        return axios.post(this.joinUrl(BASE, api.tasks, api.addTask), task).then(response => response.data);
    }
}