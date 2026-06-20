<script lang="ts">
  import { onMount } from "svelte";
  import { currentPage, PathContainers } from "../../states/page";
  import { Terminal } from '@xterm/xterm';
  import { WebLinksAddon } from '@xterm/addon-web-links';
  import "@xterm/xterm/css/xterm.css";
  import { FitAddon } from "@xterm/addon-fit";


  export let params: {
    id: string;
  };
  
  onMount(() => {
    currentPage.set(PathContainers);
  });

  

let terminalElement: HTMLElement;

onMount(() => {
  
const term = new Terminal({
  cursorBlink: true,
});

const fitAddon = new FitAddon()
term.loadAddon(fitAddon);
term.loadAddon(new WebLinksAddon());

term.open(terminalElement);

const ws = new WebSocket("ws://localhost:3000/api/containers/docker/terminal/" + params.id);

      ws.onmessage = (e) => {
        term.write(e.data);
      };

      term.onData((data) => {
        ws.send(data);
      });

       const resizeObserver = new ResizeObserver(() => {
    fitAddon.fit();
  });

  resizeObserver.observe(terminalElement);

  return () => {
    resizeObserver.disconnect();
    ws.close();
    term.dispose();
  };
})
</script>

<div class="w-full h-full overflow-hidden p-2 bg-black rounded-md">
  <div class="w-full h-full" bind:this={terminalElement}></div>
</div>