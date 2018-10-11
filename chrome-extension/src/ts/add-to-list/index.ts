import axios from 'axios'

export const addCardByCodeAndNumber = (set: string, number: string) => {
    axios.post(`http://localhost:1323/entry`, { set, number })
}
