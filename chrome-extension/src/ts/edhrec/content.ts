import { addCardByCodeAndNumber } from '@/add-to-list'

const createButtonWrapper = (code: string, number: string) => {
    const wrapper = document.createElement('div')
    wrapper.className = 'mtg-deck-tools button-wrapper'

    const addButton = document.createElement('button')
    addButton.className = 'btn btn-default'
    addButton.innerText = 'Save in list'
    addButton.onclick = () => addCardByCodeAndNumber(code, number)

    wrapper.appendChild(addButton)

    return wrapper
}

const getCodeAndNumberFromImageSrc = (card: HTMLDivElement): { code: string, number: string, ok: boolean } => {
    const img = card.getElementsByTagName('img')[0]

    if (img === undefined || img.src === '') return { code: '', number: '', ok: false }

    const params = img.src.replace('https://img.scryfall.com/cards/', '').split('/')
    if (params[2] === undefined || params[3] === undefined) return { code: '', number: '', ok: false }

    const number = params[3].split('.')[0]

    return { code: params[2], number, ok: true }
}

let processedCards: HTMLDivElement[] = []
const processCards = () => {
    const cards = document.getElementsByClassName('card') as HTMLCollectionOf<HTMLDivElement>
    for (let i = 0; i < cards.length; i++) {
        const card = cards[i]

        if (processedCards.includes(card)) continue

        const priceTags = card.getElementsByClassName('price')
        for (let j = 0; j < priceTags.length; j++) {
            card.removeChild(priceTags[j])
        }

        const codeAndNumber = getCodeAndNumberFromImageSrc(card)
        if (!codeAndNumber.ok) continue

        const buttonWrapper = createButtonWrapper(codeAndNumber.code, codeAndNumber.number)
        card.parentElement!.after(buttonWrapper)

        processedCards.push(card)
    }
}

const onScroll = (callback: () => void) => {
    let timer: number;

    window.addEventListener('scroll', function () {
        if (timer) window.clearTimeout(timer)

        timer = window.setTimeout(callback, 100)
    });
}

export default () => {
    processCards()
    onScroll(processCards)
}