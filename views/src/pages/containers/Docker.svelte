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
		<button class="{tabSelected.name === 'detail'
			? 'bg-neutral-800 border-forgelog text-forgelog'
			: 'border-neutral-700'
		} border-b-2  py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
		onclick={showDetail}>Detail</button>
		
		<button class="{tabSelected.name === 'network'
			? 'bg-neutral-800 border-forgelog text-forgelog'
			: 'border-neutral-700'
		} border-b-2  py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
		onclick={showNetwork}>Network</button>
		
		<button class="{tabSelected.name === 'environment'
			? 'bg-neutral-800 border-forgelog text-forgelog'
			: 'border-neutral-700'
		} border-b-2  py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
		onclick={() => console.log('environment')}>Environment</button>
		
		<button class="{tabSelected.name === 'config'
			? 'bg-neutral-800 border-forgelog text-forgelog'
			: 'border-neutral-700'
		} border-b-2  py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
		onclick={() => console.log('config')}>Config</button>

		<button class="{tabSelected.name === 'terminal'
			? 'bg-neutral-800 border-forgelog text-forgelog'
			: 'border-neutral-700'
		} border-b-2  py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
		onclick={showTerminal}>Terminal</button>
		</div>
	</div>

	<div class="flex-1 min-h-0">
	<tabSelected.component {...tabSelected.props} />
	</div>
</div>
