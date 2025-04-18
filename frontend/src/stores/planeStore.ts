import { writable } from 'svelte/store';
import type { Plane } from '../models/Plane';
import { round4 } from '../lib/map/utils';

const planeMap = new Map<string, Plane>();

const { subscribe, update, set } = writable<Map<string, Plane>>(planeMap);

export interface PlaneWithTrack extends Plane {
  track: [number, number][];
}

export function updatePlane(id: string, newData: Partial<Plane>) {
  update(currentMap => {
    const existing = currentMap.get(id) ?? {
      id,
      latitude: 0,
      longitude: 0,
      altitude: 0,
      heading: 0,
      speed: 0,
      timestamp: ''
    };

    if (newData.latitude)
      newData.latitude = round4(newData.latitude)
    if (newData.longitude)
      newData.longitude = round4(newData.longitude)


    const updated = { ...existing, ...newData, id };
    currentMap.set(id, updated);

    return new Map(currentMap); // trigger reactivity
  });
}

export function setPlanes(planes: Plane[]) {
  const map = new Map<string, Plane>();
  planes.forEach(plane => map.set(plane.id, plane));
  set(map);
}

export const planeStore = {
  subscribe,
  updatePlane,
  setPlanes,
};
