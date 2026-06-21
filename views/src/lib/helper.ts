export function getLastPath(): string {
  const hash = window.location.hash;
  const lastPath = hash.split('/').pop();

  return lastPath || '';
}