import type { Discussion } from "src/api";
import { writable } from "svelte/store";

export const activePost = writable<Discussion | null>(null);
