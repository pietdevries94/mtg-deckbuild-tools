import buttons from './buttons'

export default () => {
    const actions = document.getElementsByClassName('card-actions')

    for (let i = 0; i < actions.length; i++) {
        const button = <HTMLButtonElement>(actions[i].getElementsByTagName('button')[0])
        const [set, code] = button.dataset.id!.split(':')
        actions[i].innerHTML = ''
        actions[i].appendChild(buttons(set, code))
    }
}