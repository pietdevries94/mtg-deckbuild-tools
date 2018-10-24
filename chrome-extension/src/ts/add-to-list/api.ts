import axios, { AxiosError } from 'axios'

export interface CardInterface {
    scryfall_id: string
    set_code: string
    set_number: string
    name: string
    oracle_id: string
    thumbnail_url: string
    updated_at: string
}

export interface ListInterface {
    id: number
    name: string
    included_tags: string[]
}

export interface EntryInterface {
    id?: number
    scryfall_id: string
    list_id: number
    tags: string[]
}

export interface GetCardResponse {
    card: CardInterface
    entries: EntryInterface[]
}

export interface GetListsResponse {
    lists: ListInterface[]
}

export interface GetTagsResponse {
    tags: string[]
}

export interface PostListPayload {
    name: string
}

export interface PostListResponse {
    id: number
}

const client = axios.create({
    baseURL: "http://localhost:1323",
    timeout: 3000,
})

export async function getCardBySetAndNumber(set: string, number: string) {
    try {
        const response = await client.get<GetCardResponse>("/card/set-number/" + set + "/" + number)
        return response.data
    } catch (e) {
        throw (e)
    }
}

export async function getCardByName(name: string) {
    try {
        const response = await client.get<GetCardResponse>("/card/name/" + name)
        return response.data
    } catch (e) {
        throw (e)
    }
}

export function postEntry(payload: EntryInterface) {
    return client.post<void>("/entry", payload)
}

export async function getLists() {
    try {
        const response = await client.get<GetListsResponse>("/list")
        return response.data
    } catch (e) {
        throw (e)
    }
}

export async function getTags() {
    try {
        const response = await client.get<GetTagsResponse>("/tag")
        return response.data
    } catch (e) {
        throw (e)
    }
}

export async function postList(payload: PostListPayload) {
    try {
        const response = await client.post<PostListResponse>("/list", payload)
        return response.data
    } catch (e) {
        throw (e)
    }
}

export function deleteEntry(scryfallID: string, listID: number) {
    return client.delete("/entry/" + scryfallID + "/" + listID)
}