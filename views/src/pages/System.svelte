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
</script>

<div class="flex flex-col h-full min-h-0 gap-3">
	<div class="flex text-sm relative z-10">
		<span class="h-0.5 bg-neutral-700 z-9 w-full absolute bottom-0"></span>
		<div class="flex-row flex relative z-10">
			<button
				class="{tabSelected.name === 'detail'
					? 'bg-neutral-800 border-forgelog text-forgelog'
					: 'border-neutral-700'} border-b-2 py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
				onclick={showDetail}>Detail</button
			>

			<button
				class="{tabSelected.name === 'network'
					? 'bg-neutral-800 border-forgelog text-forgelog'
					: 'border-neutral-700'} border-b-2 py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
				>Network</button
			>


			<button
				class="{tabSelected.name === 'terminal'
					? 'bg-neutral-800 border-forgelog text-forgelog'
					: 'border-neutral-700'} border-b-2 py-2 px-10 cursor-pointer hover:bg-neutral-700/60"
				onclick={showTerminal}>Terminal</button
			>
		</div>
	</div>

	<div class="flex-1 min-h-0">
		<tabSelected.component {...tabSelected.props} />
	</div>
</div>
