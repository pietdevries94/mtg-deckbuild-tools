import { addCardByName } from "@/add-to-list";

const createButton = (name: string): HTMLButtonElement => {
    const button = document.createElement('button')
    button.className = 'btn btn-success mtg-deck-tools-add'
    button.onclick = () => addCardByName(name)

    const icon = document.createElement('span')
    icon.className = 'glyphicon glyphicon-plus'
    button.appendChild(icon)

    return button
}

const fillTable = (mutations: MutationRecord[]) => {
    for (let i = 0; i < mutations.length; i++) {
        if (mutations[i].target.nodeName !== 'TBODY') continue

        for (let j = 0; j < mutations[i].addedNodes.length; j++) {
            const row = mutations[i].addedNodes[j] as HTMLTableRowElement
            const tds = row.getElementsByTagName('td')
            const first = tds[0]
            const last = tds[tds.length - 1]

            if (last.getElementsByClassName('mtg-deck-tools-add').length > 0) continue

            const name = first.innerText

            last.getElementsByTagName('div')[0].appendChild(createButton(name))
        }
    }
}

export default () => {
    const table = document.getElementById('inventory-table') as HTMLTableElement
    if (table === null) return

    const observer = new MutationObserver(fillTable)
    observer.observe(table, { childList: true, subtree: true })
}