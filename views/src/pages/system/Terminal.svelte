<script lang="ts">
	import { Terminal } from '@xterm/xterm';
	import { WebLinksAddon } from '@xterm/addon-web-links';
	import '@xterm/xterm/css/xterm.css';
	import { FitAddon } from '@xterm/addon-fit';
	import { onDestroy, onMount } from 'svelte';
	import { Maximize, Minimize } from '@lucide/svelte';

	let terminalElement: HTMLElement;
  let containerElement: HTMLElement;

	let ws: WebSocket;
	let term: Terminal;
	let resizeObserver: ResizeObserver;

  let isFullscreen = $state(false);
  let isConnected = $state(false);

	function wsConnector(): WebSocket {
		return new WebSocket('ws://localhost:3000/api/system/terminal');
	}

  function reconnect() {
    term.clear();
    term.clearSelection();
    term.reset();

    ws = wsConnector();
    ws.onmessage = (e) => term.write(e.data);
    ws.onclose = () => isConnected = false;
    ws.onerror = (ev) => {
      isConnected = false;
      console.error('WebSocket error :',ev);
    }

    setTimeout(() => {
      if (ws.OPEN) {
        isConnected = true;
      }
    }, 500);
  }

  async function toggleFullscreen() {
		if (!document.fullscreenElement) {
			await containerElement.requestFullscreen();
		} else {
			await document.exitFullscreen();
		}
	}

	onMount(() => {
		// Initialize Terminal
		term = new Terminal({ cursorBlink: true });
		const fitAddon = new FitAddon();
		term.loadAddon(fitAddon);
		term.loadAddon(new WebLinksAddon());
		term.open(terminalElement);
    fitAddon.fit();

		// Open WebSocket connection
		ws = wsConnector();
    isConnected = true;

		ws.onmessage = (e) => term.write(e.data);
    ws.onclose = () => isConnected = false;
    ws.onerror = (ev) => {
      isConnected = false;
      console.error('WebSocket error :',ev);
    }

		term.onData((data) => ws.send(data));

		// Terminal resize observer
		resizeObserver = new ResizeObserver(() => fitAddon.fit());
		resizeObserver.observe(terminalElement);

    // Fullscreen listener
		const fullscreenHandler = () => {
			isFullscreen = !!document.fullscreenElement;

			// Wait browser finish fullscreen layout
			requestAnimationFrame(() => {
				fitAddon.fit();
			});
		};

    document.addEventListener('fullscreenchange', fullscreenHandler);

		return () => {
			document.removeEventListener(
				'fullscreenchange',
				fullscreenHandler
			);
		};
	});

	onDestroy(() => {
		resizeObserver?.disconnect();
		ws?.close();
		term?.dispose();
	});
</script>

<div
bind:this={containerElement}
class="w-full h-full p-2 bg-black rounded-md relative">
  <button
  onclick={toggleFullscreen}
  class="absolute top-2 right-2 bg-neutral-900 hover:bg-neutral-700
    p-2 rounded-md z-90 cursor-pointer">
    {#if isFullscreen}
			<Minimize size={16} />
		{:else}
			<Maximize size={16} />
		{/if}
  </button>
  {#if !isConnected}
  <div class="absolute top-0 left-0 bottom-0 right-0 z-90 flex items-center justify-center">
    <div class="bg-amber-400/20 text-amber-500 py-8 px-16 w-fit h-fit rounded-md flex flex-col items-center gap-1">
      <span class="text-lg">Terminal is diconnected</span>
      <button 
      onclick={reconnect}
      class="text-sm bg-amber-600 text-white hover:bg-amber-500 py-1 px-3 w-fit rounded-md cursor-pointer">Reconnect</button>
    </div>
  </div>
  {/if}
  <div class="w-full h-full" bind:this={terminalElement}></div>
</div>