import type { AxiosResponse } from "axios";
import axios from "./axios";

export interface Discussion {
	id: string;
	number: number;
	upvoteCount: number;
	answer: Answer | null;
	author: Author;
	comments: Comments;
	reactions: Comments;
	bodyHTML: string;
	title: string;
	createdAt: string;
	url: string;
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

export interface Comments {
	totalCount: number;
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
