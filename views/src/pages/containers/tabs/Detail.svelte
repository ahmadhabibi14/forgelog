<script lang="ts">
	import axios from 'axios';
	import { onMount, tick } from 'svelte';
	import '@andypf/json-viewer';

	let params: {
		containerId: string;
	} = $props();

	let containerDetail: any | {} = $state({});
	let codesJson: HTMLElement;

	let isCodeReady: boolean = $state(false);

	onMount(async () => {
		try {
			const { data } = await axios.get(`/api/container/${params.containerId}`);

			containerDetail = data;

			isCodeReady = true;
			await tick();
		} catch (err) {
			console.error(err);
		}
	});
</script>

<div class="flex w-full h-full rounded-md relative">
	<div class="overflow-auto scrollbar-custom h-full w-full relative">
		<div class=" flex flex-col gap-1">
			<div class="bg-neutral-800 p-2 rounded-md">chyyy</div>
			<div class="overflow-x-hidden h-auto w-full rounded-md">
				{#if isCodeReady}
					<andypf-json-viewer
						theme="default-dark"
						indent="2"
						expanded="true"
						show-data-types="true"
						show-toolbar="false"
						expand-icon-type="arrow"
						show-copy="true"
						show-size="true"
						expand-empty="false"
					>
						{JSON.stringify(containerDetail, null, 2)}
					</andypf-json-viewer>
				{/if}
			</div>
		</div>
	</div>
</div>
