import type { AxiosResponse } from "axios";
import axios from "./axios";

export const authenticate = async (code: string): Promise<AxiosResponse> => {
	return await axios.post("/auth/", {}, { params: { code } });
};

export const checkAuth = async (): Promise<AxiosResponse> => {
	return await axios.get("/auth/");
};

export const logout = async (): Promise<AxiosResponse> => {
	return await axios.delete("/auth/");
};
