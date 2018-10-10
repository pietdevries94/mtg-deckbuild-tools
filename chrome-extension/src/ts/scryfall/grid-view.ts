import buttons from './buttons'

const cleanup = () => {
    const oldElements = document.getElementsByClassName('mtg-deck-tools')
    while (oldElements[0]) {
        oldElements[0].parentNode!.removeChild(oldElements[0]);
    }
}

interface gridInfo { ok: boolean, cardsOnPage: number, cardsInRow: number, cards?: HTMLCollectionOf<HTMLAnchorElement> }

const getGridInfo = (): gridInfo => {
    const result: gridInfo = { ok: false, cardsOnPage: 0, cardsInRow: 0 }

    const cardGridInners = document.getElementsByClassName('card-grid-inner')
    if (cardGridInners.length === 0) {
        return result
    }

    const cardGridInner = <HTMLDivElement>cardGridInners[0]

    result.cards = <HTMLCollectionOf<HTMLAnchorElement>>cardGridInner.getElementsByClassName('card-grid-item')
    if (result.cards.length === 0) {
        return result
    }

    result.cardsOnPage = result.cards.length

    if (window.innerWidth >= 920) {
        result.cardsInRow = 4
    } else if (window.innerWidth >= 660) {
        result.cardsInRow = 3
    } else {
        result.cardsInRow = 2
    }

    result.ok = true
    return result
}

const getCodeAndNumberFromAnchor = (card: HTMLAnchorElement): { code: string, number: number } => {
    const params = card.href.replace('https://scryfall.com/card/', '').split('/')

    return { code: params[0], number: Number(params[1]) }
}

const getPlaceHolder = (): HTMLAnchorElement => {
    const placeholder = document.createElement('a')
    placeholder.className = 'mtg-deck-tools card-grid-item'
    const image = document.createElement('img')
    image.src = chrome.runtime.getURL('images/card-placeholder.jpeg')
    image.className = 'card ons border-black'
    placeholder.appendChild(image)
    return placeholder
}

export default () => {
    cleanup()

    const info = getGridInfo()

    let buttonWrappers: HTMLDivElement[] = []

    for (let i = 0; i < info.cardsOnPage; i++) {
        // add a wrapper around the anchor
        const card = info.cards![i]

        const cn = getCodeAndNumberFromAnchor(card)

        buttonWrappers.push(buttons(cn.code, cn.number))

        const cardNumber = i + 1
        if (cardNumber % info.cardsInRow === 0 || cardNumber === info.cardsOnPage) {
            buttonWrappers.reverse().forEach(button => {
                card.after(button)
            })
            buttonWrappers = []
        }


        // add placeholders for last row
        if (cardNumber === info.cardsOnPage && (cardNumber % info.cardsInRow) > 0) {
            const numberOfPlaceHolders = info.cardsInRow - (cardNumber % info.cardsInRow)
            for (let j = 0; j < numberOfPlaceHolders; j++) {
                card.after(getPlaceHolder())
            }
        }
    }
}
