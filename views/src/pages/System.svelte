<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { currentPage, PathSystem } from '../states/page';
	import Detail from './system/Detail.svelte';
	import Terminal from './system/Terminal.svelte';

	onMount(() => {
		currentPage.set(PathSystem);
	});

	let tabSelected: {
		name: string;
    component: any;
		props: Object;
	} = $state({
		name: 'detail',
    component: Detail,
    props: {}
	});

  function showDetail() {
    tabSelected = {
      name: 'detail',
      component: Detail,
      props: {}
    };
  }

  function showTerminal() {
    tabSelected = {
      name: 'terminal',
      component: Terminal,
      props: {}
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
			func: () => console.log('network')
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
