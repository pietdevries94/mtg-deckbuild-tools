import Axios from "axios";

const importInventory = async (user: string) => {
    const res = await Axios.get(`/users/${user}/inventory/?fmt=csv`)
    Axios.post("http://localhost:1323/api/inventory", res.data)
}

const createButton = (user: string): HTMLButtonElement => {
    const button = document.createElement('button')
    button.className = 'btn btn-sm btn-danger btn-block mtg-deck-tools-add'
    button.innerText = 'Import inventory'
    button.onclick = () => {
        importInventory(user)
        button.disabled = true
        button.className = 'btn btn-sm btn-block mtg-deck-tools-add'
    }

    return button
}

export default () => {
    const path = window.location.pathname.split('/')
    if (path[1] !== 'users' || path[3] !== 'inventory') return

    const btn = createButton(path[2])

    const col = document.getElementsByClassName('well')[2]!.getElementsByClassName('col-lg-3')[2]!.getElementsByTagName('p')[0]!
    col.innerHTML = ''
    col.appendChild(btn)

    // importInventory('kwartel')
}