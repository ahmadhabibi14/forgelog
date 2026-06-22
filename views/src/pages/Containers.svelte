<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';
	import type { DockerContainer } from '../types/docker';
	import * as Table from '$lib/components/ui/table/index.js';
	import { formatUnixTimestamp } from '$lib/formatter';
	import { currentPage, PathContainers } from '../states/page';
	import { Funnel, Plus, Search } from '@lucide/svelte';
	import * as Select from '$lib/components/ui/select/index.js';
	import { Input } from '$lib/components/ui/input/index.js';

	let inputSearch: HTMLInputElement | any = $state(null);
	let inputSearchText: string = $state('');

	onMount(() => {
		window.addEventListener('keydown', (e) => {
			console.log('key', e.key);
			if (e.key === '/') {
				e.preventDefault();

				inputSearchText = '';
				inputSearch.focus();
			}
		});
		currentPage.set(PathContainers);
	});

	let ContainersDocker: DockerContainer[] = $state([]);

	async function getContainersDocker() {
		await axios
			.post('/api/containers')
			.then((res) => {
				console.log(res.data);
				ContainersDocker = res.data.Items || [];
			})
			.catch((err) => {
				console.error(err);
			});
	}

	onMount(async () => {
		await getContainersDocker();
	});

	const columns = [
		{ value: 'name', label: 'Name' },
		{ value: 'container_id', label: 'Container ID' },
		{ value: 'image', label: 'Image' }
	];

	let columnToFilter = $state('');

	const triggerContent = $derived(columns.find((f) => f.value === columnToFilter)?.label ?? 'Name');
</script>

<div class="flex flex-col gap-2 h-full min-h-0">
	<div class="flex flex-row justify-between items-center">
		<div class="flex flex-row gap-3 items-center p-2">
			<button>
				<Funnel size="16" />
			</button>
			<Select.Root type="single" name="favoriteFruit" bind:value={columnToFilter}>
				<Select.Trigger class="w-35">
					{triggerContent}
				</Select.Trigger>
				<Select.Content>
					<Select.Group>
						<Select.Label>Search by</Select.Label>
						{#each columns as col (col.value)}
							<Select.Item value={col.value} label={col.label}>
								{col.label}
							</Select.Item>
						{/each}
					</Select.Group>
				</Select.Content>
			</Select.Root>

			<div class="relative">
				<Input
					bind:ref={inputSearch}
					bind:value={inputSearchText}
					type="text"
					placeholder="Search..."
					class="max-w-3xl w-100 pr-9"
				/>
				<div
					class="absolute top-1.5 right-1 bg-neutral-800 h-6 w-6 border border-neutral-700 rounded flex items-center justify-center"
				>
					<span class="text-xs font-bold">/</span>
				</div>
			</div>
		</div>
		<div>
			<button
				class="bg-forgelog hover:bg-forgelog-light text-neutral-950 py-1.5 px-3 rounded-md
		flex flex-row gap-2 items-center text-sm cursor-pointer"
			>
				<Plus size="17" />
				<span>Add New</span>
			</button>
		</div>
	</div>
	<div class="flex-1 overflow-y-auto scrollbar-custom">
		<Table.Root class="sticky-table relative">
			<Table.Header class="top-0 z-10 bg-neutral-950 w-full">
				<Table.Row class="">
					<Table.Head class="sticky top-0 w-25">Container ID</Table.Head>
					<Table.Head>Image</Table.Head>
					<Table.Head>Name</Table.Head>
					<Table.Head>Created</Table.Head>
					<Table.Head class="text-end">Status</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each ContainersDocker || [] as docker}
					{@const createdTime = new Date(docker.Created * 1000)}
					<Table.Row
						class="group cursor-pointer"
						onclick={() => (window.location.href = `#/containers/${docker.Id}`)}
					>
						<Table.Cell class="group-hover:underline group-hover:text-blue-500 font-medium"
							>{docker.Id.slice(0, 12)}</Table.Cell
						>
						<Table.Cell class="group-hover:underline group-hover:text-blue-500"
							>{docker.Image}</Table.Cell
						>
						<Table.Cell class="group-hover:underline group-hover:text-blue-500"
							>{docker.Names[0]}</Table.Cell
						>
						<Table.Cell class="group-hover:underline group-hover:text-blue-500">
							<time datetime={createdTime.toString()}>{formatUnixTimestamp(docker.Created)}</time>
						</Table.Cell>
						<Table.Cell class="text-end">
							<span class="{docker.State} text-xs py-1 px-3 rounded-full">
								{docker.Status}
							</span>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
</div>

<style lang="postcss">
	@reference "../app.css";

	.exited {
		@apply bg-red-300/10 text-red-500;
	}

	.running {
		@apply bg-forgelog-light/15 text-forgelog;
	}

	:global(.sticky-table) {
		overflow: visible !important;
	}
</style>
