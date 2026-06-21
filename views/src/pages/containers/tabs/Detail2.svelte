<script lang="ts">
	import axios from "axios";
	import { onMount, tick } from "svelte";
  import Highlight from 'svelte-highlight/Highlight.svelte';
  import LineNumbers from  'svelte-highlight/LineNumbers.svelte'; 
  import json from 'svelte-highlight/languages/json';
  import style from 'svelte-highlight/styles/github-dark';

  let params: {
    containerId: string;
  } = $props();

  let containerDetail: any|{} = $state({});
  let codesJson: HTMLElement;

  let isCodeReady: boolean = $state(false);

  onMount(async () => {
    try {
    const { data } = await axios.get(
      `/api/containers/docker/detail/${params.containerId}`
    );

    containerDetail = data;

    isCodeReady = true;
    await tick();
  } catch (err) {
    console.error(err);
  }
});
  
</script>


<svelte:head>
  {@html style}
</svelte:head>

<div class="flex w-full h-full">
<div class="overflow-y-scroll overflow-x-hidden max-h-full w-full rounded-md">
  {#if isCodeReady}
    <Highlight language={json} code={JSON.stringify(containerDetail, null, 2)} let:highlighted>
      <LineNumbers {highlighted}/>
    </Highlight>
  {/if}
</div>
</div>

<style>
  :global(pre) {
  white-space: pre-wrap !important;
  word-break: break-word;
  overflow-wrap: anywhere;
  overflow-x: hidden;
  max-width: 100%;
}

:global(pre code.hljs) {
  white-space: inherit;
}
</style>