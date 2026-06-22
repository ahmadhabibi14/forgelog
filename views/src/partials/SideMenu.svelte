<script lang="ts">
	import {
		Columns3Cog,
		Container,
		LayoutDashboard,
		FileClock,
		type LucideProps,
		Settings,
		GlobeLock,

		MonitorCog

	} from '@lucide/svelte';
	import type { Component } from 'svelte';
	import { currentPage } from '../states/page';

	const navmenu: {
		path: string;
		icon: Component<LucideProps, {}, ''>;
		label: string;
	}[] = [
		{
			path: '#/',
			icon: LayoutDashboard,
			label: 'Dashboard'
		},
		{
			path: '#/containers',
			icon: Container,
			label: 'Containers'
		},
		{
			path: '#/logs',
			icon: FileClock,
			label: 'Logging'
		},
		{
			path: '#/configs',
			icon: Columns3Cog,
			label: 'Configs'
		},
		{
			path: '#/networks',
			icon: GlobeLock,
			label: 'Networks'
		},
		{
			path: '#/system',
			icon: MonitorCog,
			label: 'System'
		},
		{
			path: '#/settings',
			icon: Settings,
			label: 'Settings'
		}
	];
</script>

<aside class="w-3xs h-full p-4 flex flex-col gap-1">
	{#each navmenu as nav (nav.path)}
		<a
			href={nav.path}
			onclick={() => ($currentPage = nav.path)}
			class="flex flex-row gap-3 items-center py-2 px-2 hover:bg-neutral-900
      rounded-md active:ring-0 focus:ring-0 focus:border-none active:border-none
        focus:bg-neutral-900 focus:outline-0
        {$currentPage == nav.path ? 'bg-neutral-900 text-forgelog' : ''}"
		>
			<svelte:component this={nav.icon} size="18" />
			<span>{nav.label}</span>
		</a>
	{/each}
</aside>
