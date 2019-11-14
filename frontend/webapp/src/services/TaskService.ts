import {injectable} from 'tsyringe'
import BaseService from './BaseService';
import { Task } from '@/models/task';
import axios from 'axios'
import { BASE, api } from './api';

@injectable()
export default class TaskService extends BaseService
{
    tasks(): Promise<Array<Task>>
    {
        return axios.get(this.joinUrl(BASE, api.tasks)).then(response => response.data);
    }
    moveTask(task: Task): Promise<any>
    {
        return axios.put(this.joinUrl(BASE, api.moveTask), {id: task.id}).then(response => response.data);
    }
    addTask(task: Task): Promise<any>
    {
        return axios.post(this.joinUrl(BASE, api.addTask), task).then(response => response.data);
    }
}