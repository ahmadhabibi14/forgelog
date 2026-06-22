export function getLastPath(): string {
  const hash = window.location.hash;
  const lastPath = hash.split('/').pop();

  return lastPath || '';
}

export function websocketConnector(path: string): WebSocket {
  if (!path.startsWith('/')) {
    path = '/' + path;
  }

  if (import.meta.env.PROD) {
    const host = window.location.host;
    if (host.includes('localhost')) {
      return new WebSocket(`ws://${host}${path}`);
    } 
    return new WebSocket(`wss://${host}${path}`);
  } else {
    return new WebSocket(`ws://${import.meta.env.VITE_API_PLAIN_BASE_URL}${path}`);
  }
}