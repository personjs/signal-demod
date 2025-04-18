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

    if (newData.type !== "Airborne Position") {
      newData.latitude = round4(existing.latitude);
      newData.longitude = round4(existing.longitude);
      newData.speed = existing.speed;
    }


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
