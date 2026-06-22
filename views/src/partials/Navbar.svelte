<script lang="ts">
	import { ArrowUpDown, Cpu, HardDrive, MemoryStick, TextAlignJustify } from '@lucide/svelte';
	import type { SystemStat } from '../types/stats';
	import { onMount } from 'svelte';

  onMount(() => {
    console.log('HOST:', window.location.host)
    console.log('ENV', import.meta.env);
  })

  let systemStats: SystemStat = $state({
    bandwidth: '',
    cpu: '',
    disk: '',
    memory: ''
  });

  const es = new EventSource(`${import.meta.env.VITE_API_BASE_URL}/api/system/stats`);
	es.addEventListener("stats", (e) => {
		const data = JSON.parse(e.data) || {};
    systemStats = data;
	});
</script>

<div
	class="fixed h-16 top-0 right-0 left-0 bg-neutral-950 w-full flex flex-row justify-between px-4 items-center"
>
	<div class="flex flex-row items-center gap-4">
		<button
			class="flex justify-center items-center p-2
        hover:bg-neutral-900 active:ring-0 focus:ring-0 focus:border-none active:border-none
        focus:bg-neutral-900 focus:outline-0 rounded-lg cursor-pointer"
		>
			<TextAlignJustify size="20" />
		</button>
		<div>
			<img src="/main-logo.svg" alt="Forgelog" class="h-6" />
		</div>
	</div>

  <div class="flex flex-row gap-2 items-center text-sm">
    <div class="flex flex-row gap-2 items-center py-1 px-3 bg-violet-400/20 text-violet-500 rounded-full">
      <MemoryStick size="17" />
      <span>{systemStats.memory}</span>
    </div>
    <div class="flex flex-row gap-2 items-center py-1 px-3 bg-rose-400/20 text-rose-500 rounded-full">
      <Cpu size="17" />
      <span>{systemStats.cpu}</span>
    </div>
    <div class="flex flex-row gap-2 items-center py-1 px-3 bg-amber-400/20 text-amber-500 rounded-full">
      <HardDrive size="17" />
      <span>{systemStats.disk}</span>
    </div>
    <div class="flex flex-row gap-2 items-center py-1 px-3 bg-forgelog-light/15 text-forgelog rounded-full">
      <ArrowUpDown size="17" />
      <span>{systemStats.bandwidth}</span>
    </div>
  </div>
</div>
