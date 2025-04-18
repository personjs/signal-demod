import { planeStore } from '../stores/planeStore';
import type { Plane } from '../models/Plane';

let socket: WebSocket;

export function initPlaneSocket(url: string) {
  socket = new WebSocket(url);

  socket.onopen = () => console.log('[WebSocket] Connected');
  socket.onmessage = handleMessage;
  socket.onclose = () => console.warn('[WebSocket] Disconnected');
  socket.onerror = (err) => {
    console.error('[WebSocket] Error:', err);
    socket.close();
  };
}

function handleMessage(event: MessageEvent) {
  try {
    const data: Plane = JSON.parse(event.data);
    if (data.type === "Airborne Position" && data.latitude !== 0 && data.longitude !== 0) {
      planeStore.updatePlane(data.id, data);
    }
  } catch (err) {
    console.error('[WebSocket] Failed to parse message:', err);
  }
}

export function closePlaneSocket() {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.close();
    console.log('[WebSocket] Closed by client');
  }
}
