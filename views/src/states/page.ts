import { writable, type Writable } from 'svelte/store';

export const PathDashboard: string = '#/';
export const PathContainers: string = '#/containers';

export const currentPage: Writable<string> = writable(PathDashboard);
