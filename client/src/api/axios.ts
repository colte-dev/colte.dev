import { dev } from "$app/env";
import axios from "axios";

const baseURL = dev ? "http://localhost:3000/api" : "/api";
const client = axios.create({
	baseURL,
	validateStatus: () => true,
	withCredentials: true,
});

export default client;
