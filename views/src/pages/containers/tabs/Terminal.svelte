<script lang="ts">
  import { Terminal } from '@xterm/xterm';
	import { WebLinksAddon } from '@xterm/addon-web-links';
	import '@xterm/xterm/css/xterm.css';
	import { FitAddon } from '@xterm/addon-fit';
	import { onDestroy, onMount } from 'svelte';

  export let containerId: string = '';

	let terminalElement: HTMLElement;

  let ws: WebSocket;
  let term: Terminal;
  let resizeObserver: ResizeObserver;

  function wsConnector(): WebSocket {
    return new WebSocket('ws://localhost:3000/api/containers/docker/terminal/' + containerId);
  }

  onMount(() => {
    // Initialize Terminal
		term = new Terminal({ cursorBlink: true });
		const fitAddon = new FitAddon();
		term.loadAddon(fitAddon);
		term.loadAddon(new WebLinksAddon());
		term.open(terminalElement);

    // Open WebSocket connection
		ws = wsConnector();
		ws.onmessage = (e) => term.write(e.data);
		term.onData(data => ws.send(data));

    // Terminal resize observer
		resizeObserver = new ResizeObserver(() => fitAddon.fit());
		resizeObserver.observe(terminalElement);
	});

  onDestroy(() => {
    resizeObserver.disconnect();
    ws.close();
    term.dispose();
  });
</script>

<div class="w-full h-full p-2 bg-black rounded-md">
	<div
		class="w-full h-full"
		bind:this={terminalElement}
	></div>
</div>