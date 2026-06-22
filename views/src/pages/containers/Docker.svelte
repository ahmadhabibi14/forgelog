<script lang="ts">
	import { onMount } from 'svelte';
	import { currentPage, PathContainers } from '../../states/page';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb/index.js';
	import { House } from '@lucide/svelte';
	import Terminal from './tabs/Terminal.svelte';
	import Detail from './tabs/Detail.svelte';
	import { getLastPath } from '$lib/helper';
	import Network from './tabs/Network.svelte';

	let containerId: string = getLastPath();

	onMount(() => {
		currentPage.set(PathContainers);
	});

	
	let tabSelected: {
		name: string;
		component: any;
		props: Object
	} = $state({
		name: 'detail',
		component: Detail,
		props: {
			containerId: containerId
		}
	});

	function showDetail() {
		tabSelected = {
		name: 'detail',
		component: Detail,
		props: {
			containerId: containerId
		}
	};
	}

	function showNetwork() {
		tabSelected = {
		name: 'network',
		component: Network,
		props: {
			containerId: containerId
		}
	};
	}

	function showTerminal() {
		tabSelected = {
		name: 'terminal',
		component: Terminal,
		props: {
			containerId: containerId
		}
	};
	}

	const tabsMenu: {
		name: string;
		label: string;
		func: () => void;
	}[] = [
		{
			name: 'detail',
			label: 'Detail',
			func: showDetail
		},
		{
			name: 'network',
			label: 'Network',
			func: showNetwork
		},
		{
			name: 'config',
			label: 'Config',
			func: () => console.log('config')
		},
		{
			name: 'terminal',
			label: 'Terminal',
			func: showTerminal
		},
	]
</script>

<div class="flex flex-col h-full min-h-0 gap-3">
	<Breadcrumb.Root class="text-sm! py-2 pl-2">
		<Breadcrumb.List>
			<Breadcrumb.Item>
				<Breadcrumb.Link href="#/" class="flex flex-row gap-2 items-center">
					<House size="16" />
					<span>Dashboard</span>
				</Breadcrumb.Link>
			</Breadcrumb.Item>
			<Breadcrumb.Separator />
			<Breadcrumb.Item>
				<Breadcrumb.Link href="#/containers">Containers</Breadcrumb.Link>
			</Breadcrumb.Item>
			<Breadcrumb.Separator />
			<Breadcrumb.Item>
				<Breadcrumb.Link href="#/containers">Docker</Breadcrumb.Link>
			</Breadcrumb.Item>
			<Breadcrumb.Separator />
			<Breadcrumb.Item>
				<Breadcrumb.Page>{containerId.slice(0,12)}</Breadcrumb.Page>
			</Breadcrumb.Item>
		</Breadcrumb.List>
	</Breadcrumb.Root>
	<div class="flex text-sm relative z-10">
		<span class="h-0.5 bg-neutral-700 z-9 w-full absolute bottom-0"></span>
		<div class="flex-row flex relative z-10">
		{#each tabsMenu as tab (tab.name)}
			<button class="{tabSelected.name === tab.name
				? 'border-forgelog text-forgelog'
				: 'border-neutral-700 hover:bg-neutral-700/60'
			} border-b-2  py-2 px-6 cursor-pointer 
			flex items-center justify-center"
			onclick={tab.func}>
				<span class="{tabSelected.name === tab.name ? 'bg-neutral-800' : ''}
				w-fit h-fit py-1 px-4 rounded-md">{tab.label}</span>
			</button>
		{/each}
		</div>
	</div>

	<div class="flex-1 min-h-0">
	<tabSelected.component {...tabSelected.props} />
	</div>
</div>
