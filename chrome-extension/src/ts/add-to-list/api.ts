import axios, { AxiosPromise } from 'axios'

export interface CardInterface {
    scryfall_id: string
    set_code: string
    set_number: string
    name: string
    oracle_id: string
    thumbnail_url: string
    updated_at: string
}

export interface PostEntryPayload {
    scryfall_id: string
}

const client = axios.create({
    baseURL: "http://localhost:1323"
})

export async function getCardBySetAndNumber(set: string, number: string) {
    const response = await client.get<CardInterface>("/card/set-number/" + set + "/" + number)
    return response.data
}

export function postEntry(payload: PostEntryPayload) {
    return client.post<void>("/entry", payload)
}