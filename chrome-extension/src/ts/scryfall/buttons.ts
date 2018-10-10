import { addCardByCodeAndNumber } from '@/add-to-list'

const createWrapper = (): HTMLDivElement => {
    const buttonWrapper = document.createElement('div')
    buttonWrapper.className = "mtg-deck-tools custom-card-grid-item"
    return buttonWrapper
}

const createButton = (text: string): HTMLButtonElement => {
    const button = document.createElement('button')
    button.type = 'button'
    button.className = 'button-n'
    button.innerText = text
    return button
}

export default (code: string, number: number): HTMLDivElement => {
    const buttonWrapper = createWrapper()
    const addButton = createButton('Save in list')

    addButton.onclick = () => addCardByCodeAndNumber(code, number)

    buttonWrapper.appendChild(addButton)
    return buttonWrapper
}