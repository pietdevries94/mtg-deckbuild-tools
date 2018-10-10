export default () => {
    const oldElements = document.getElementsByClassName('mtg-deck-tools')
    while (oldElements[0]) {
        oldElements[0].parentNode!.removeChild(oldElements[0]);
    }

    const cardGridInners = document.getElementsByClassName('card-grid-inner')
    if (cardGridInners.length === 0) {
        return
    }

    const cardGridInner = <HTMLDivElement>cardGridInners[0]

    const cardGridItems = cardGridInner.getElementsByClassName('card-grid-item')
    if (cardGridItems.length === 0) {
        return
    }

    const cardsOnPage = cardGridItems.length

    let cardsInRow: number = 2
    if (window.innerWidth >= 920) {
        cardsInRow = 4
    } else if (window.innerWidth >= 660) {
        cardsInRow = 3
    }

    let buttonWrappers: HTMLDivElement[] = []
    for (let i = 0; i < cardsOnPage; i++) {
        // add a wrapper around the anchor
        const card = <HTMLAnchorElement>cardGridItems[i]

        const buttonWrapper = document.createElement('div')
        buttonWrapper.className = "mtg-deck-tools custom-card-grid-item"
        const button = document.createElement('button')
        button.type = 'button'
        button.className = 'button-n'
        button.innerText = 'Save in list'
        buttonWrapper.appendChild(button)
        buttonWrappers.push(buttonWrapper)

        const cardNumber = i + 1
        if (cardNumber % cardsInRow === 0 || cardNumber === cardsOnPage) {
            buttonWrappers.reverse().forEach(button => {
                card.after(button)
            })
            buttonWrappers = []
        }

        // add placeholders for last row
        if (cardNumber === cardsOnPage && (cardNumber % cardsInRow) > 0) {
            const numberOfPlaceHolders = cardsInRow - (cardNumber % cardsInRow)
            for (let j = 0; j < numberOfPlaceHolders; j++) {
                const placeholder = document.createElement('a')
                placeholder.className = 'mtg-deck-tools card-grid-item'
                const image = document.createElement('img')
                image.src = chrome.runtime.getURL('images/card-placeholder.jpeg')
                image.className = 'card ons border-black'
                placeholder.appendChild(image)
                card.after(placeholder)
            }
        }
    }
}
