import type { AxiosResponse } from "axios";
import axios from "./axios";

export interface Discussion {
	id: string;
	number: number;
	upvoteCount: number;
	answer: Answer | null;
	author: Author;
	comments: Nodes<Comment>;
	reactions: Nodes;
	bodyHTML: string;
	title: string;
	createdAt: string;
	url: string;
}

export interface Comment {
	id: string;
	author: Author;
	bodyHTML: string;
	createdAt: string;
	reactions: Nodes;
	upvoteCount: number;
}

export interface Answer {
	id: string;
	url: string;
}

export interface Author {
	login: string;
	url: string;
	avatarUrl: string;
}

export interface Nodes<T = undefined> {
	totalCount: number;
	nodes: T[];
}

export const getDiscussionByNumber = async (number: number): Promise<AxiosResponse<Discussion>> => {
	const response = await axios.get(`/discussions/${number.toString()}`);
	return response;
};

export const getDiscussions = async (): Promise<AxiosResponse<Discussion[]>> => {
	const response = await axios.get("/discussions/");
	return response;
};

export const getUserDiscussions = async (
	username: string
): Promise<AxiosResponse<Discussion[]>> => {
	const response = await axios.get("/discussions/", { params: { username } });
	return response;
};
